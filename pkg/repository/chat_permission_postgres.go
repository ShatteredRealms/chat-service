package repository

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/srospan"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type chatChannelPermissionPostgresRepository struct {
	gormdb *gorm.DB
}

func NewChatChannelPermissionPostgresRepository(db *gorm.DB) ChatChannelPermissionRepository {
	db.AutoMigrate(&chat.ChannelPermission{})
	return &chatChannelPermissionPostgresRepository{gormdb: db}
}

// AddForCharacter implements ChatChannelPermissionRepository.
func (r *chatChannelPermissionPostgresRepository) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)
	tx := r.db(ctx).Begin()
	for _, channelId := range channelIds {
		result := tx.Save(&chat.ChannelPermission{
			ChannelId:   channelId,
			CharacterId: characterId,
		})
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

// RemForCharacter implements ChatChannelPermissionRepository.
func (r *chatChannelPermissionPostgresRepository) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)
	tx := r.db(ctx).Begin()
	for _, channelId := range channelIds {
		result := tx.Delete(&chat.ChannelPermission{}, "channel_id = ? AND character_id = ?", channelId, characterId)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

// GetForCharacter implements ChatChannelPermissionRepository.
func (r *chatChannelPermissionPostgresRepository) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	tagCharacter(ctx, characterId)

	var channels chat.Channels
	result := r.db(ctx).
		Model(&chat.Channel{}).
		Joins("JOIN channel_permissions ON channels.id = channel_permissions.channel_id").
		Where("channel_permissions.character_id = ?", characterId).
		Find(&channels)
	return &channels, result.Error
}

// HasAccess implements ChatChannelPermissionRepository.
func (r *chatChannelPermissionPostgresRepository) HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error) {
	tagCharacter(ctx, characterId)

	result := r.db(ctx).
		Where("channel_id = ? AND character_id = ?", channelId, channelId).
		First(&chat.ChannelPermission{})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// SaveForCharacter implements ChatChannelPermissionRepository.
func (r *chatChannelPermissionPostgresRepository) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)

	tx := r.db(ctx).Begin()
	result := tx.Delete(&chat.ChannelPermission{}, "character_id = ?", characterId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	for _, channelId := range channelIds {
		result = tx.Create(&chat.ChannelPermission{
			ChannelId:   channelId,
			CharacterId: characterId,
		})
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	return tx.Commit().Error
}

func (r *chatChannelPermissionPostgresRepository) db(ctx context.Context) *gorm.DB {
	return r.gormdb.WithContext(ctx)
}

func tagCharacter(ctx context.Context, characterId string) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		srospan.TargetCharacterId(characterId),
	)
}
