package service

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/google/uuid"
)

type ChatChannelService interface {
	GetAll(ctx context.Context) (*chat.Channels, error)
	GetById(ctx context.Context, id *uuid.UUID) (*chat.Channel, error)

	Create(ctx context.Context, name string, dimensionId string) (*chat.Channel, error)
	Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error)
	Delete(ctx context.Context, channelId *uuid.UUID) error
}

type chatChannelService struct {
	repo repository.ChatChannelRepository
}

func NewChatChannelService(repo repository.ChatChannelRepository) ChatChannelService {
	return &chatChannelService{
		repo: repo,
	}
}

// Create implements ChatChannelService.
func (c *chatChannelService) Create(ctx context.Context, name string, dimensionId string) (*chat.Channel, error) {
	channel := &chat.Channel{
		Name:        name,
		DimensionId: dimensionId,
	}
	return c.repo.Create(ctx, channel)
}

// Delete implements ChatChannelService.
func (c *chatChannelService) Delete(ctx context.Context, channelId *uuid.UUID) error {
	return c.repo.Delete(ctx, channelId)
}

// GetAll implements ChatChannelService.
func (c *chatChannelService) GetAll(ctx context.Context) (*chat.Channels, error) {
	return c.repo.GetAll(ctx)
}

// GetById implements ChatChannelService.
func (c *chatChannelService) GetById(ctx context.Context, id *uuid.UUID) (*chat.Channel, error) {
	return c.repo.GetById(ctx, id)
}

// Save implements ChatChannelService.
func (c *chatChannelService) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	return c.repo.Save(ctx, channel)
}
