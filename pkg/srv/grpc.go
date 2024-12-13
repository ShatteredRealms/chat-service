package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/pb"
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

	RoleChatChannelBan = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("chat.channel.ban"),
		Description: gocloak.StringP("Allows banning and unbanning characters from chat channels"),
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
	channel, character, err := s.validateChannelPermissions(server.Context(), request.ChannelId, request.CharacterId, chat.PermissionRead)
	if err != nil {
		return err
	}

	dimensionId := channel.DimensionId
	if dimensionId == nil {
		dimensionId = &character.DimensionId
	}
	chatMessages, err := s.Context.ChatService.ReceiveChannelMessages(server.Context(), &channel.Id, dimensionId, &character.Id)

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

	character, err := s.validateCharacterOwner(server.Context(), request.Id, claims.Subject)
	if err != nil {
		return err
	}

	chatMessages, err := s.Context.ChatService.ReceiveDirectMessages(server.Context(), &character.Id, &character.Id)

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
					Errorf("sending grpc message for character '%s' by '%s': %v",
						request.Id, claims.Subject, err)
				return err
			}
		}
	}
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

	err = s.Context.ChatChannelService.Delete(ctx, request.Id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatDelete, err)
		return nil, ErrChatDelete
	}

	return &emptypb.Empty{}, nil
}

// EditChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) EditChatChannel(ctx context.Context, request *pb.UpdateChatChannelRequest) (*emptypb.Empty, error) {
	_, err := s.validateRole(ctx, RoleChatManagement)
	if err != nil {
		return nil, err
	}

	_, err = s.Context.ChatChannelService.Update(ctx, request)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatEdit, err)
		return nil, status.Errorf(codes.InvalidArgument, ErrChatEdit.Error())
	}

	return &emptypb.Empty{}, nil
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

	subjectId, err := uuid.Parse(claim.Subject)
	if err != nil {
		return nil, errors.New("invalid owner id")
	}
	if character.OwnerId != subjectId &&
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

	channel, err := s.Context.ChatChannelService.GetById(ctx, request.Id)
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
	channel, character, err := s.validateChannelPermissions(ctx, request.ChannelId, request.ChatMessage.SenderCharacterId, chat.PermissionReadSend)
	if err != nil {
		return nil, err
	}

	dimensionId := channel.DimensionId
	if dimensionId == nil {
		dimensionId = &character.DimensionId
	}
	err = s.Context.ChatService.SendChannelMessage(ctx, &channel.Id, dimensionId, &chat.Message{
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

	senderCharacter, err := s.validateCharacterOwner(ctx, request.ChatMessage.SenderCharacterId, claims.Subject)
	if err != nil {
		return nil, err
	}

	targetCharacter, err := s.getCharacter(ctx, request.CharacterId)
	if err != nil {
		return nil, err
	}

	if targetCharacter.DimensionId != senderCharacter.DimensionId {
		return nil, status.Errorf(codes.NotFound, "character not found")
	}

	err = s.Context.ChatService.SendDirectMessage(ctx, &targetCharacter.Id, &chat.Message{
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

// BanCharacterFromChatChannel implements pb.ChatServiceServer.
func (s *chatServiceServer) BanCharacterFromChatChannel(ctx context.Context, request *pb.BanRequest) (*emptypb.Empty, error) {
	_, err := s.validateRole(ctx, RoleChatChannelBan)
	if err != nil {
		return nil, err
	}

	channelId, err := uuid.Parse(request.ChannelId)
	if err != nil {
		return nil, ErrChatIdInvalid
	}

	_, err = s.getCharacter(ctx, request.CharacterId)
	if err != nil {
		return nil, err
	}

	err = s.Context.ChatChannelPermissionService.BanCharacter(ctx, request.CharacterId, &channelId, request.Duration)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("error banning character '%s' from channel '%s': %v", request.CharacterId, request.ChannelId, err)
		return nil, status.Error(codes.Internal, ErrChatPermissionSet.Error())
	}

	return &emptypb.Empty{}, nil
}
