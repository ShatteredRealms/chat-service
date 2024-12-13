package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type chatHistoryPgxRepository struct {
	conn *pgxpool.Pool
}

func NewChatHistoryPgxRepository(migrater *PgxMigrater) ChatHistoryRepository {
	return &chatHistoryPgxRepository{
		conn: migrater.conn,
	}
}

// AddMessage implements ChatHistoryRepository.
func (c *chatHistoryPgxRepository) AddMessage(ctx context.Context, channelId *uuid.UUID, dimensionId *uuid.UUID, message *chat.Message) error {
	if message == nil {
		return ErrEmptyMessage
	}

	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	senerId, err := uuid.Parse(message.SenderCharacterId)
	if err != nil {
		return err
	}
	ct, err := tx.Exec(ctx, "INSERT INTO chat_messages (channel_id, sender_id, content, created_at) VALUES ($1, $2, $3, $4)",
		channelId, senerId, message.Content, time.Now().UTC())
	if err != nil {
		return err
	}
	if ct.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return tx.Commit(ctx)
}

// GetMessages implements ChatHistoryRepository.
func (c *chatHistoryPgxRepository) GetMessages(ctx context.Context, channelId string, limit *int, offset *int, before *time.Time, after *time.Time, sender *string) (msgs *chat.Messages, maxOffset int, err error) {
	panic("unimplemented")
}
