package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
)

// TODO: Think about implementing GetMessageCount, that takes the same arguments as
// GetMessages, except limit and offset, and returns the count of messages
// that would be returned by GetMessages.

type ChatLogRepository interface {
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
		channelId *uuid.UUID,
		limit *uint,
		offset *uint,
		before *time.Time,
		after *time.Time,
		sender *uuid.UUID,
	) (msgs *chat.MessageLogs, total uint, err error)
}
