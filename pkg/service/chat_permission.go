package service

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/google/uuid"
)

type ChatChannelPermissionService interface {
	GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error)
	SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error)
}

type chatChannelPermissionService struct {
	repo repository.ChatChannelPermissionRepository
}

func NewChatChannelPermissionService(repo repository.ChatChannelPermissionRepository) ChatChannelPermissionService {
	return &chatChannelPermissionService{
		repo: repo,
	}
}

// GetForCharacter implements ChatChannelPermissionService.
func (c *chatChannelPermissionService) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	return c.repo.GetForCharacter(ctx, characterId)
}

// SaveForCharacter implements ChatChannelPermissionService.
func (c *chatChannelPermissionService) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.SaveForCharacter(ctx, characterId, channelIds)
}

// HasAccess implements ChatChannelPermissionService.
func (c *chatChannelPermissionService) HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error) {
	return c.repo.HasAccess(ctx, channelId, characterId)
}

// AddForCharacter implements ChatChannelPermissionService.
func (c *chatChannelPermissionService) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.AddForCharacter(ctx, characterId, channelIds)
}

// RemForCharacter implements ChatChannelPermissionService.
func (c *chatChannelPermissionService) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.RemForCharacter(ctx, characterId, channelIds)
}
