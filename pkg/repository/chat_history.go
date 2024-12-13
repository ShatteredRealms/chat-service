package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
)

type ChatHistoryRepository interface {
	AddMessage(
		ctx context.Context,
		channelId *uuid.UUID,
		dimensionId *uuid.UUID,
		message *chat.Message,
	) error

	// GetMessages retrieves messages from a channel. The channel Id may be a
	// chat channel or a character
	GetMessages(
		ctx context.Context,
		channelId string,
		limit *int,
		offset *int,
		before *time.Time,
		after *time.Time,
		sender *string,
	) (msgs *chat.Messages, maxOffset int, err error)
}
