package srv

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/character/characterbus"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	minLevel chat.ChannelPermissionLevel,
) (*chat.Channel, *characterbus.Character, error) {
	claims, err := s.validateRole(ctx, RoleChatChannelUse)
	if err != nil {
		return nil, nil, err
	}

	channel, err := s.getChatChannel(ctx, channelId)
	if err != nil {
		return nil, nil, err
	}

	character, err := s.validateCharacterOwner(ctx, characterId, claims.Subject)
	if err != nil {
		return nil, nil, err
	}

	// Validate character has perssions to access channel
	level, err := s.Context.ChatChannelPermissionService.GetAccessLevel(ctx, &channel.Id, characterId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrChatPermissionGet, err)
		return nil, nil, status.Error(codes.Internal, ErrChatPermissionGet.Error())
	}
	if !(channel.Public && level > chat.PermissionPermBan) && (!channel.Public && level < minLevel) {
		log.Logger.WithContext(ctx).
			Warnf("user '%s' and character '%s' tried accessing '%s' but has no permissions",
				claims.Subject, characterId, channelId)
		return nil, nil, srv.ErrPermissionDenied
	}
	return channel, character, nil
}

func (s *chatServiceServer) getChatChannel(ctx context.Context, channelId string) (*chat.Channel, error) {
	channel, err := s.Context.ChatChannelService.GetById(ctx, channelId)
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

func (s *chatServiceServer) validateCharacterOwner(ctx context.Context, characterId string, ownerId string) (*characterbus.Character, error) {
	// Validate sender owns character
	character, err := s.getCharacter(ctx, characterId)
	if err != nil {
		return nil, err
	}

	oId, err := uuid.Parse(ownerId)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("unable to process request: %v", err)
		return nil, status.Error(codes.Internal, "unable to process request")
	}

	if character.OwnerId != oId {
		log.Logger.WithContext(ctx).
			Warnf("user '%s' tried chat acesss for character '%s' with owner '%s'",
				ownerId, character.Id, character.OwnerId)
		return nil, srv.ErrPermissionDenied
	}

	return character, nil
}
