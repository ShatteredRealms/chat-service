package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type chatChannelPgxRepository struct {
	conn *pgxpool.Pool
}

func NewChatChannelPgxRepository(migrater *PgxMigrater) ChatChannelRepository {
	return &chatChannelPgxRepository{
		conn: migrater.conn,
	}
}

// Create implements ChatChannelRepository.
func (p *chatChannelPgxRepository) Create(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "INSERT INTO chat_channels (name, dimension_id) VALUES ($1, $2) RETURNING *", channel.Name, channel.DimensionId)
	if err != nil {
		return nil, err
	}

	outChannel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[chat.Channel])
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &outChannel, nil
}

// Delete implements ChatChannelRepository.
func (p *chatChannelPgxRepository) Delete(ctx context.Context, channelId *uuid.UUID) error {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}

	ct, err := tx.Exec(ctx, "DELETE FROM chat_channels WHERE id = $1", channelId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrDoesNotExist
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements ChatChannelRepository.
func (p *chatChannelPgxRepository) GetAll(ctx context.Context) (*chat.Channels, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT id, name, dimension_id, created_at, updated_at FROM chat_channels WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	channels := make(chat.Channels, 0)
	for rows.Next() {
		channel := &chat.Channel{}
		err = rows.Scan(&channel.Id, &channel.Name, &channel.DimensionId, &channel.CreatedAt, &channel.UpdatedAt)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &channels, nil
}

// GetById implements ChatChannelRepository.
func (p *chatChannelPgxRepository) GetById(ctx context.Context, id *uuid.UUID) (channel *chat.Channel, _ error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	outChannel, err := p.queryById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return outChannel, nil
}

// Save implements ChatChannelRepository.
func (p *chatChannelPgxRepository) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	ct, err := tx.Exec(ctx,
		"UPDATE chat_channels SET name = $2, dimension_id = $3, updated_at = $4 WHERE id = $1",
		channel.Id, channel.Name, channel.DimensionId, time.Now(),
	)
	if ct.RowsAffected() == 0 {
		return nil, ErrDoesNotExist
	}

	outChannel, err := p.queryById(ctx, tx, &channel.Id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return outChannel, nil
}

func (p *chatChannelPgxRepository) queryById(ctx context.Context, tx pgx.Tx, id *uuid.UUID) (*chat.Channel, error) {
	outChannel := &chat.Channel{}
	err := tx.QueryRow(ctx, "SELECT id, name, dimension_id, created_at, updated_at FROM chat_channels WHERE id = $1", id).
		Scan(&outChannel.Id, &outChannel.Name, &outChannel.DimensionId, &outChannel.CreatedAt, &outChannel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return outChannel, nil
}
