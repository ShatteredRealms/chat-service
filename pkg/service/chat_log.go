package service

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/google/uuid"
)

type ChatLogService interface {
	AddMessage(
		ctx context.Context,
		channelId *uuid.UUID,
		dimensionId *uuid.UUID,
		message *chat.Message,
	) error

	GetMessages(ctx context.Context, request *pb.ChatLogRequest) (msgs *chat.MessageLogs, total uint, err error)
}

type chatLogService struct {
	repo repository.ChatLogRepository
}

func NewChatLogService(repo repository.ChatLogRepository) ChatLogService {
	return &chatLogService{
		repo: repo,
	}
}

// GetMessages implements ChatLogService.
func (c *chatLogService) GetMessages(ctx context.Context, request *pb.ChatLogRequest) (*chat.MessageLogs, uint, error) {
	var channelId uuid.UUID
	var limit *uint
	var offset *uint
	var before *time.Time
	var after *time.Time
	var sender *uuid.UUID
	var err error

	channelId, err = uuid.Parse(request.GetChannelId())
	if err != nil {
		return nil, 0, err
	}

	if request.GetLimit() > 0 {
		limit = new(uint)
		*limit = uint(request.GetLimit())
	}
	if request.GetOffset() > 0 {
		offset = new(uint)
		*offset = uint(request.GetOffset())
	}
	if request.GetBefore() > 0 {
		before = &time.Time{}
		*before = time.Unix(request.GetBefore(), 0)
	}
	if request.GetAfter() > 0 {
		after = &time.Time{}
		*after = time.Unix(request.GetAfter(), 0)
	}
	if request.GetSenderId() != "" {
		id, err := uuid.Parse(request.GetSenderId())
		if err != nil {
			return nil, 0, err
		}
		sender = &id
	}

	return c.repo.GetMessages(ctx, &channelId, limit, offset, before, after, sender)
}

// AddMessage implements ChatLogService.
func (c *chatLogService) AddMessage(ctx context.Context, channelId *uuid.UUID, dimensionId *uuid.UUID, message *chat.Message) error {
	return c.repo.AddMessage(ctx, channelId, dimensionId, message)
}
