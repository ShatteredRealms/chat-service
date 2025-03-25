package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type chatChannelPgxRepository struct {
	conn *pgxpool.Pool
}

func NewChatChannelPgxRepository(migrater *repository.PgxMigrater) ChatChannelRepository {
	return &chatChannelPgxRepository{
		conn: migrater.Conn,
	}
}

// Create implements ChatChannelRepository.
func (p *chatChannelPgxRepository) Create(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "INSERT INTO chat_channels (name, dimension_id, public) VALUES ($1, $2, $3) RETURNING *", channel.Name, channel.DimensionId, channel.Public)
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
func (p *chatChannelPgxRepository) Delete(ctx context.Context, channelId *uuid.UUID) (*chat.Channel, error) {
	if channelId == nil {
		return nil, ErrNilId
	}

	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"UPDATE chat_channels SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *",
		channelId)
	outChannel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[chat.Channel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return outChannel, tx.Commit(ctx)

}

// GetAll implements ChatChannelRepository.
func (p *chatChannelPgxRepository) GetAll(ctx context.Context) (*chat.Channels, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT id, name, dimension_id, public, created_at, updated_at FROM chat_channels WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	channels := make(chat.Channels, 0)
	for rows.Next() {
		channel := &chat.Channel{}
		err = rows.Scan(&channel.Id, &channel.Name, &channel.DimensionId, &channel.Public, &channel.CreatedAt, &channel.UpdatedAt)
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
		"UPDATE chat_channels SET name = $2, dimension_id = $3, public = $4 WHERE id = $1",
		channel.Id, channel.Name, channel.DimensionId, channel.Public,
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

// Update implements ChatChannelRepository.
func (p *chatChannelPgxRepository) Update(ctx context.Context, request *UpdateRequest) (*chat.Channel, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]any)
	if request.Name != nil {
		updates["name"] = *request.Name
	}
	if request.DimensionId != nil {
		if *request.DimensionId == "" || *request.DimensionId == "null" {
			updates["dimension_id"] = nil
			log.Logger.Infof("dimension_id: %s", updates["dimension_id"])
		} else {
			updates["dimension_id"] = *request.DimensionId
		}
	}
	if request.Public != nil {
		if *request.Public {
			updates["public"] = "true"
		} else {
			updates["public"] = "false"
		}
	}

	if len(updates) == 0 {
		return nil, ErrNoUpdates
	}

	builder := strings.Builder{}
	builder.WriteString("UPDATE chat_channels SET ")
	vals := make([]any, 0, len(updates)+1)
	vals = append(vals, request.ChannelId)
	argNum := 2
	for column, value := range updates {
		builder.WriteString(fmt.Sprintf("%s = $%d", column, argNum))
		if argNum < len(updates)+1 {
			builder.WriteString(", ")
		}
		argNum++
		vals = append(vals, value)
	}
	builder.WriteString(" WHERE id = $1 RETURNING *")
	log.Logger.Infof("builder: %s", builder.String())
	log.Logger.Infof("vals: %s", vals)

	rows, _ := tx.Query(ctx, builder.String(), vals...)
	outChannel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[chat.Channel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDoesNotExist
		}
		return nil, err
	}

	err = outChannel.Validate()
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
	err := tx.QueryRow(ctx, "SELECT id, name, dimension_id, public, created_at, updated_at FROM chat_channels WHERE id = $1", id).
		Scan(&outChannel.Id, &outChannel.Name, &outChannel.DimensionId, &outChannel.Public, &outChannel.CreatedAt, &outChannel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return outChannel, nil
}
