package repository

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
)

type ChatChannelPermissionRepository interface {
	GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error)
	SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error)
}
