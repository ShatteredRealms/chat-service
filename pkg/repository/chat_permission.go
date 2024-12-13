package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
)

type ChatChannelPermissionRepository interface {
	GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error)
	SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	GetAccessLevel(ctx context.Context, channelId *uuid.UUID, characterId string) (chat.ChannelPermissionLevel, error)
	BanCharacter(ctx context.Context, characterId string, channelId *uuid.UUID, until *time.Time) error
}
