package service

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/google/uuid"
)

type ChatChannelPermissionService interface {
	GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error)
	SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error
	GetAccessLevel(ctx context.Context, channelId *uuid.UUID, characterId string) (chat.ChannelPermissionLevel, error)
	BanCharacter(ctx context.Context, characterId string, channelId *uuid.UUID, durationSec int64) error
}

type ccpService struct {
	repo repository.ChatChannelPermissionRepository
}

func NewChatChannelPermissionService(repo repository.ChatChannelPermissionRepository) ChatChannelPermissionService {
	return &ccpService{
		repo: repo,
	}
}

// GetAccessLevel implements ChatChannelPermissionService.
func (c *ccpService) GetAccessLevel(
	ctx context.Context,
	channelId *uuid.UUID,
	characterId string,
) (chat.ChannelPermissionLevel, error) {
	return c.repo.GetAccessLevel(ctx, channelId, characterId)
}

// GetForCharacter implements ChatChannelPermissionService.
func (c *ccpService) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	return c.repo.GetForCharacter(ctx, characterId)
}

// SaveForCharacter implements ChatChannelPermissionService.
func (c *ccpService) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.SaveForCharacter(ctx, characterId, channelIds)
}

// AddForCharacter implements ChatChannelPermissionService.
func (c *ccpService) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.AddForCharacter(ctx, characterId, channelIds)
}

// RemForCharacter implements ChatChannelPermissionService.
func (c *ccpService) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	return c.repo.RemForCharacter(ctx, characterId, channelIds)
}

// BanCharacter implements ChatChannelPermissionService.
func (c *ccpService) BanCharacter(ctx context.Context, characterId string, channelId *uuid.UUID, durationSec int64) error {
	if durationSec == -1 {
		return c.repo.BanCharacter(ctx, characterId, channelId, nil)
	}

	if durationSec == 0 {
		return c.repo.BanCharacter(ctx, characterId, channelId, &time.Time{})
	}

	if durationSec < 0 {
		return repository.ErrInvalidBanDuration
	}

	time := time.Now().UTC().Add(time.Duration(durationSec) * time.Second)
	return c.repo.BanCharacter(ctx, characterId, channelId, &time)
}
