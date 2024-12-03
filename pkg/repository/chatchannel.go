package repository

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
)

type ChatChannelRepository interface {
	GetAll(ctx context.Context) (*chat.Channels, error)
	GetById(ctx context.Context, id *uuid.UUID) (*chat.Channel, error)

	Create(ctx context.Context, channel *chat.Channel) (*chat.Channel, error)
	Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error)
	Delete(ctx context.Context, channelId *uuid.UUID) error
}