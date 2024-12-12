package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/character/characterbus"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ChatRoles = make([]*gocloak.Role, 0)

	RoleChatManagement = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.manage"),
		Description: gocloak.StringP("Allows creating, reading and editing chats channels"),
	}, &ChatRoles)

	RoleChatPermissionsManagement = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.permissions.manage"),
		Description: gocloak.StringP("Allows managing chat channel permissions for all characters"),
	}, &ChatRoles)

	RoleChatChannelUse = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.channel.use"),
		Description: gocloak.StringP("Allows reading and sending messages in chat channels as yourself"),
	}, &ChatRoles)

	RoleChatDirectUse = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.direct.use"),
		Description: gocloak.StringP("Allows reading and sending direct chat messages to others as yourself"),
	}, &ChatRoles)

	RoleChatBusReset = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.bus.reset"),
		Description: gocloak.StringP("Allows resetting the chat bus to reprocess all messages"),
	}, &ChatRoles)
)

var (
	ErrChatDoesNotExist = errors.New("ChS-Ch-00")
	ErrChatCreate       = errors.New("ChS-Ch-01")
	ErrChatDelete       = errors.New("ChS-Ch-02")
	ErrChatGet          = errors.New("ChS-Ch-03")
	ErrChatEdit         = errors.New("ChS-Ch-04")
	ErrChatIdInvalid    = errors.New("ChS-Ch-07")
	ErrChatExpired      = errors.New("ChS-Ch-08")

	ErrChatPermissionGet = errors.New("ChS-ChP-01")
	ErrChatPermissionSet = errors.New("error saving permissions for chat channel")

	ErrChatSend = errors.New("ChS-ChM-01")

	ErrDimensionNotExist = errors.New("ChS-D-01")
	ErrDimensionLookup   = errors.New("ChS-D-02")

	ErrCharacterNotExist = errors.New("ChS-Ca-01")
	ErrCharacterLookup   = errors.New("ChS-Ca-02")
)

type chatServiceServer struct {
	pb.UnimplementedChatServiceServer
	Context *ChatContext
}

func NewChatServiceServer(ctx context.Context, chatCtx *ChatContext) (pb.ChatServiceServer, error) {
	err := chatCtx.CreateRoles(ctx, &ChatRoles)
	if err != nil {
		return nil, err
	}
	return &chatServiceServer{Context: chatCtx}, nil
}

// ConnectChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) ConnectChatChannel(request *pb.ConnectChatChannelRequest, server grpc.ServerStreamingServer[pb.ChatMessage]) error {
	claims, id, err := s.validateChannelPermissions(server.Context(), request.ChannelId, request.CharacterId)
	if err != nil {
		return err
	}

	chatMessages, err := s.Context.ChatService.ReceiveChannelMessages(server.Context(), id, claims.Subject)

	for {
		select {
		case <-server.Context().Done():
			return nil
		case msg, ok := <-chatMessages:
			if !ok {
				return fmt.Errorf("receiver shutdown")
			}
			err = server.Send(&pb.ChatMessage{
				SenderCharacterId: msg.SenderCharacterId,
				Content:           msg.Content,
			})
			if err != nil {
				log.Logger.WithContext(server.Context()).
					Errorf("error sending chat message on channel '%s' for character '%s': %v",
						request.ChannelId, request.CharacterId, err)
				return err
			}
		}
	}
}

// ConnectDirectMessages implements pb.ChatServiceServer.
func (s *chatServiceServer) ConnectDirectMessages(request *commonpb.TargetId, server grpc.ServerStreamingServer[pb.ChatMessage]) error {
	// Validate general authorization
	claims, err := s.validateRole(server.Context(), RoleChatChannelUse)
	if err != nil {
		return err
	}

	err = s.validateCharacterOwner(server.Context(), request.Id, claims.Subject)
	if err != nil {
		return err
	}

	// Receive messages
	for server.Context().Err() == nil {
		msg, err := s.Context.ChatService.ReceiveDirectMessage(server.Context(), request.Id, claims.Subject)
		if err != nil {
			log.Logger.WithContext(server.Context()).
				Errorf("receiving message for character '%s' by '%s': %v",
					request.Id, claims.Subject, err)
			return err
		}

		err = server.Send(&pb.ChatMessage{
			SenderCharacterId: msg.SenderCharacterId,
			Content:           msg.Content,
		})
		if err != nil {
			log.Logger.WithContext(server.Context()).
				Errorf("sending grpc message for character '%s' by '%s': %v",
					request.Id, claims.Subject, err)
			return err
		}
	}

	return nil
}

// CreateChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) CreateChatChannel(ctx context.Context, request *pb.CreateChatChannelMessage) (*emptypb.Empty, error) {
	_, err := s.validateRole(ctx, RoleChatManagement)
	if err != nil {
		return nil, err
	}

	dimension, err := s.Context.DimensionService.GetDimensionById(ctx, request.DimensionId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionLookup, err)
		return nil, status.Error(codes.Internal, ErrDimensionLookup.Error())
	}
	if dimension == nil {
		return nil, status.Error(codes.InvalidArgument, ErrDimensionNotExist.Error())
	}

	_, err = s.Context.ChatChannelService.Create(ctx, request.Name, request.DimensionId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatCreate, err)
		return nil, ErrChatCreate
	}

	return &emptypb.Empty{}, nil
}

// DeleteChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) DeleteChatChannel(ctx context.Context, request *commonpb.TargetId) (*emptypb.Empty, error) {
	_, err := s.validateRole(ctx, RoleChatManagement)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, ErrChatIdInvalid
	}

	err = s.Context.ChatChannelService.Delete(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatDelete, err)
		return nil, ErrChatDelete
	}

	return &emptypb.Empty{}, nil
}

// EditChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) EditChatChannel(context.Context, *pb.UpdateChatChannelRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}

// GetAuthorizedChatChannels implements pb.ChatServiceServer.
func (s *chatServiceServer) GetAuthorizedChatChannels(ctx context.Context, request *commonpb.TargetId) (*pb.ChatChannels, error) {
	claim, err := s.validateRole(ctx, RoleChatChannelUse)
	if err != nil {
		return nil, err
	}

	character, err := s.getCharacter(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if character.OwnerId != claim.Subject &&
		!claim.HasResourceRole(RoleChatPermissionsManagement, s.Context.Config.Keycloak.ClientId) {
		log.Logger.WithContext(ctx).Warnf("user '%s' tried accessing chat channels for character '%s' with owner '%s'",
			claim.Subject, character.Id, character.OwnerId)
		return nil, srv.ErrPermissionDenied
	}

	// Get authorized channels
	channels, err := s.Context.ChatChannelPermissionService.GetForCharacter(ctx, request.Id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatPermissionGet, err)
		return nil, ErrChatPermissionGet
	}

	return channels.ToPb(), nil
}

// GetChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) GetChatChannel(ctx context.Context, request *commonpb.TargetId) (*pb.ChatChannel, error) {
	_, err := s.validateRole(ctx, RoleChatManagement)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, ErrChatIdInvalid
	}

	channel, err := s.Context.ChatChannelService.GetById(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatGet, err)
		return nil, ErrChatGet
	}

	if channel == nil {
		return nil, ErrChatDoesNotExist
	}

	return channel.ToPb(), nil
}

// GetChatChannels implements pb.ChatServiceServer.
func (s *chatServiceServer) GetChatChannels(ctx context.Context, _ *emptypb.Empty) (*pb.ChatChannels, error) {
	_, err := s.validateRole(ctx, RoleChatManagement)
	if err != nil {
		return nil, err
	}

	channels, err := s.Context.ChatChannelService.GetAll(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatGet, err)
		return nil, ErrChatGet
	}

	return channels.ToPb(), nil
}

// SendChatChannelMessage implements pb.ChatServiceServer.
func (s *chatServiceServer) SendChatChannelMessage(ctx context.Context, request *pb.SendChatChannelMessageRequest) (*emptypb.Empty, error) {
	_, id, err := s.validateChannelPermissions(ctx, request.ChannelId, request.ChatMessage.SenderCharacterId)
	if err != nil {
		return nil, err
	}

	err = s.Context.ChatService.SendChannelMessage(ctx, id, &chat.Message{
		SenderCharacterId: request.ChatMessage.SenderCharacterId,
		Content:           request.ChatMessage.Content,
	})

	if err != nil {
		log.Logger.WithContext(ctx).Errorf("error sending message: %v", err)
		return nil, status.Errorf(codes.Internal, ErrChatSend.Error())
	}

	return &emptypb.Empty{}, nil
}

// SendDirectMessage implements pb.ChatServiceServer.
func (s *chatServiceServer) SendDirectMessage(
	ctx context.Context,
	request *pb.SendDirectMessageRequest,
) (*emptypb.Empty, error) {
	claims, err := s.validateRole(ctx, RoleChatDirectUse)
	if err != nil {
		return nil, err
	}

	err = s.validateCharacterOwner(ctx, request.ChatMessage.SenderCharacterId, claims.Subject)
	if err != nil {
		return nil, err
	}

	err = s.Context.ChatService.SendDirectMessage(ctx, request.CharacterId, &chat.Message{
		SenderCharacterId: request.ChatMessage.SenderCharacterId,
		Content:           request.ChatMessage.Content,
	})
	if err != nil {
		log.Logger.WithContext(ctx).
			Errorf("%v: sending direct message from '%s' to '%s': %v",
				ErrChatSend.Error(), request.ChatMessage.SenderCharacterId, request.CharacterId, err)
		return nil, status.Errorf(codes.Internal, ErrChatSend.Error())
	}

	return &emptypb.Empty{}, nil
}

// SetCharacterChatChannelAuth implements pb.ChatServiceServer.
func (s *chatServiceServer) SetCharacterChatChannelAuth(
	ctx context.Context,
	request *pb.RequestSetCharacterSetChatChannelAuth,
) (*emptypb.Empty, error) {
	ids, err := s.channelAuthRequestParse(ctx, request.CharacterId, request.Ids)
	if err != nil {
		return nil, err
	}

	err = s.Context.ChatChannelPermissionService.SaveForCharacter(ctx, request.CharacterId, ids)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatPermissionSet, err)
		return nil, status.Error(codes.Internal, ErrChatPermissionSet.Error())
	}

	return &emptypb.Empty{}, nil
}

// UpdateCharacterChatChannelAuth implements pb.ChatServiceServer.
func (s *chatServiceServer) UpdateCharacterChatChannelAuth(
	ctx context.Context,
	request *pb.RequestUpdateCharacterSetChatChannelAuth,
) (*emptypb.Empty, error) {
	ids, err := s.channelAuthRequestParse(ctx, request.CharacterId, request.Ids)
	if err != nil {
		return nil, err
	}

	if request.Add {
		err = s.Context.ChatChannelPermissionService.AddForCharacter(ctx, request.CharacterId, ids)
	} else {
		err = s.Context.ChatChannelPermissionService.RemForCharacter(ctx, request.CharacterId, ids)
	}
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatPermissionSet, err)
		return nil, status.Error(codes.Internal, ErrChatPermissionSet.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *chatServiceServer) channelAuthRequestParse(ctx context.Context, charcterId string, stringIds []string) ([]*uuid.UUID, error) {
	_, err := s.validateRole(ctx, RoleChatPermissionsManagement)
	if err != nil {
		return nil, err
	}

	_, err = s.getCharacter(ctx, charcterId)
	if err != nil {
		return nil, err
	}

	ids, err := util.ParseUUIDs(stringIds)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatIdInvalid, err)
		return nil, status.Error(codes.InvalidArgument, ErrChatIdInvalid.Error())
	}

	return ids, nil
}

func (s *chatServiceServer) validateChannelPermissions(
	ctx context.Context,
	channelId, characterId string,
) (*auth.SROClaims, *uuid.UUID, error) {
	claims, err := s.validateRole(ctx, RoleChatChannelUse)
	if err != nil {
		return nil, nil, err
	}

	channel, err := s.getChatChannel(ctx, channelId)
	if err != nil {
		return nil, nil, err
	}

	err = s.validateCharacterOwner(ctx, characterId, claims.Subject)
	if err != nil {
		return nil, nil, err
	}

	// Validate character has perssions to access channel
	ok, err := s.Context.ChatChannelPermissionService.HasAccess(ctx, &channel.Id, characterId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatPermissionGet, err)
		return nil, nil, status.Error(codes.Internal, ErrChatPermissionGet.Error())
	}

	if !ok {
		log.Logger.WithContext(ctx).
			Warnf("user '%s' and character '%s' tried accessing '%s' but has no permissions",
				claims.Subject, characterId, channelId)
		return nil, nil, srv.ErrPermissionDenied
	}

	return claims, &channel.Id, nil
}

func (s *chatServiceServer) getChatChannel(ctx context.Context, channelId string) (*chat.Channel, error) {
	id, err := uuid.Parse(channelId)
	if err != nil {
		return nil, ErrChatIdInvalid
	}

	channel, err := s.Context.ChatChannelService.GetById(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("error getting chat channel: %v", err)
		return nil, status.Error(codes.Internal, ErrChatGet.Error())
	}
	if channel == nil {
		return nil, status.Error(codes.InvalidArgument, ErrChatDoesNotExist.Error())
	}

	return channel, nil
}

func (s *chatServiceServer) getCharacter(ctx context.Context, characterId string) (*characterbus.Character, error) {
	character, err := s.Context.CharacterService.GetCharacterById(ctx, characterId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("error getting character: %v", err)
		return nil, status.Error(codes.Internal, ErrCharacterLookup.Error())
	}
	if character == nil {
		return nil, status.Error(codes.InvalidArgument, ErrCharacterNotExist.Error())
	}
	return character, nil
}

func (s *chatServiceServer) validateCharacterOwner(ctx context.Context, characterId string, ownerId string) error {
	// Validate sender owns character
	character, err := s.getCharacter(ctx, characterId)
	if err != nil {
		return err
	}

	if character.OwnerId != ownerId {
		log.Logger.WithContext(ctx).
			Warnf("user '%s' tried chat acesss for character '%s' with owner '%s'",
				ownerId, character.Id, character.OwnerId)
		return srv.ErrPermissionDenied
	}

	return nil
}
