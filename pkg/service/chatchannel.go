package service

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/google/uuid"
)

type ChatChannelService interface {
	GetAll(ctx context.Context) (*chat.Channels, error)
	GetById(ctx context.Context, id string) (*chat.Channel, error)

	Create(ctx context.Context, name string, dimensionId string) (*chat.Channel, error)
	Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error)
	Update(ctx context.Context, pbRequest *pb.UpdateChatChannelRequest) (*chat.Channel, error)
	Delete(ctx context.Context, channelId string) (*chat.Channel, error)
}

type chatChannelService struct {
	repo repository.ChatChannelRepository
}

// Update implements ChatChannelService.
func (c *chatChannelService) Update(ctx context.Context, pbRequest *pb.UpdateChatChannelRequest) (*chat.Channel, error) {
	id, err := uuid.Parse(pbRequest.ChannelId)
	if err != nil {
		return nil, err
	}

	request := &repository.UpdateRequest{
		ChannelId: &id,
	}
	if pbRequest.OptionalName != nil {
		name := pbRequest.GetName()
		request.Name = &name
	}
	if pbRequest.OptionalDimension != nil {
		dimensionId := pbRequest.GetDimension()
		request.DimensionId = &dimensionId
	}

	return c.repo.Update(ctx, request)
}

func NewChatChannelService(repo repository.ChatChannelRepository) ChatChannelService {
	return &chatChannelService{
		repo: repo,
	}
}

// Create implements ChatChannelService.
func (c *chatChannelService) Create(ctx context.Context, name string, dimensionId string) (*chat.Channel, error) {
	id, err := uuid.Parse(dimensionId)
	if err != nil {
		return nil, err
	}

	channel := &chat.Channel{
		Name:        name,
		DimensionId: &id,
	}
	err = channel.Validate()
	if err != nil {
		return nil, err
	}

	return c.repo.Create(ctx, channel)
}

// Delete implements ChatChannelService.
func (c *chatChannelService) Delete(ctx context.Context, channelId string) (*chat.Channel, error) {
	id, err := uuid.Parse(channelId)
	if err != nil {
		return nil, err
	}

	return c.repo.Delete(ctx, &id)
}

// GetAll implements ChatChannelService.
func (c *chatChannelService) GetAll(ctx context.Context) (*chat.Channels, error) {
	return c.repo.GetAll(ctx)
}

// GetById implements ChatChannelService.
func (c *chatChannelService) GetById(ctx context.Context, channelId string) (*chat.Channel, error) {
	id, err := uuid.Parse(channelId)
	if err != nil {
		return nil, err
	}
	return c.repo.GetById(ctx, &id)
}

// Save implements ChatChannelService.
func (c *chatChannelService) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	return c.repo.Save(ctx, channel)
}
