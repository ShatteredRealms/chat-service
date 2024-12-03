package repository

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type chatChannelPostgresRepository struct {
	gormdb *gorm.DB
}

func NewChatChannelPostgresRepository(db *gorm.DB) ChatChannelRepository {
	db.AutoMigrate(&chat.Channel{})
	return &chatChannelPostgresRepository{gormdb: db}
}

// Create implements ChatChannelRepository.
func (p *chatChannelPostgresRepository) Create(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	return channel, p.db(ctx).Create(&channel).Error
}

// Delete implements ChatChannelRepository.
func (p *chatChannelPostgresRepository) Delete(ctx context.Context, channelId *uuid.UUID) error {
	return p.db(ctx).Delete(&chat.Channel{}, "id = ?", channelId).Error
}

// GetAll implements ChatChannelRepository.
func (p *chatChannelPostgresRepository) GetAll(ctx context.Context) (channels *chat.Channels, _ error) {
	*channels = make(chat.Channels, 0)
	return channels, p.db(ctx).Find(&channels).Error
}

// GetById implements ChatChannelRepository.
func (p *chatChannelPostgresRepository) GetById(ctx context.Context, id *uuid.UUID) (channel *chat.Channel, _ error) {
	result := p.db(ctx).First(&channel, "id = ?", id)
	return channel, result.Error
}

// Save implements ChatChannelRepository.
func (p *chatChannelPostgresRepository) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	return channel, p.db(ctx).Save(&channel).Error
}

func (p *chatChannelPostgresRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}
