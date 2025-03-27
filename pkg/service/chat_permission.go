package service

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/ShatteredRealms/go-common-service/pkg/common"
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
	id, err := uuid.Parse(characterId)
	if err != nil {
		return chat.PermissionNone, common.ErrInvalidId
	}
	return c.repo.GetAccessLevel(ctx, channelId, &id)
}

// GetForCharacter implements ChatChannelPermissionService.
func (c *ccpService) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return nil, common.ErrInvalidId
	}
	return c.repo.GetForCharacter(ctx, &id)
}

// SaveForCharacter implements ChatChannelPermissionService.
func (c *ccpService) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return common.ErrInvalidId
	}
	return c.repo.SaveForCharacter(ctx, &id, channelIds)
}

// AddForCharacter implements ChatChannelPermissionService.
func (c *ccpService) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return common.ErrInvalidId
	}
	return c.repo.AddForCharacter(ctx, &id, channelIds)
}

// RemForCharacter implements ChatChannelPermissionService.
func (c *ccpService) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return common.ErrInvalidId
	}
	return c.repo.RemForCharacter(ctx, &id, channelIds)
}

// BanCharacter implements ChatChannelPermissionService.
func (c *ccpService) BanCharacter(ctx context.Context, characterId string, channelId *uuid.UUID, durationSec int64) error {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return common.ErrInvalidId
	}
	if durationSec == 0 {
		return c.repo.BanCharacter(ctx, &id, channelId, nil)
	}

	if durationSec < 0 {
		return c.repo.BanCharacter(ctx, &id, channelId, &chat.PermBanTime)
	}

	time := time.Now().UTC().Add(time.Duration(durationSec) * time.Second)
	return c.repo.BanCharacter(ctx, &id, channelId, &time)
}
