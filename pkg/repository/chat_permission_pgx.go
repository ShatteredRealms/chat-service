package repository

import (
	"context"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ccpPgxRepo struct {
	conn *pgxpool.Pool
}

func NewChatChannelPermissionPgxRepository(migrater *PgxMigrater) ChatChannelPermissionRepository {
	return &ccpPgxRepo{
		conn: migrater.conn,
	}
}

// AddForCharacter implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)
	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	err = r.addPermissionsForUser(ctx, tx, channelIds, characterId)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// RemForCharacter implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)
	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	for _, channelId := range channelIds {
		_, err := tx.Exec(ctx, "DELETE FROM chat_channel_permissions WHERE chat_channel_id = $1 AND character_id = $2", channelId, characterId)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

// GetForCharacter implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	tagCharacter(ctx, characterId)
	channels := make(chat.Channels, 0)

	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return &channels, err
	}

	rows, err := tx.Query(
		ctx,
		`SELECT 
			chat_channels.id, 
			chat_channels.name,
			chat_channels.dimension_id,
			chat_channels.created_at,
			chat_channels.updated_at
		FROM chat_channels
		JOIN chat_channel_permissions 
			ON chat_channels.id = chat_channel_permissions.chat_channel_id
		WHERE (
			chat_channel_permissions.character_id = $1
			AND chat_channels.deleted_at IS NULL
		)`,
		characterId,
	)
	if err != nil {
		return &channels, err
	}

	for rows.Next() {
		channel := &chat.Channel{}
		err = rows.Scan(
			&channel.Id,
			&channel.Name,
			&channel.DimensionId,
			&channel.CreatedAt,
			&channel.UpdatedAt,
		)
		if err != nil {
			return &channels, err
		}
		channels = append(channels, channel)
	}

	return &channels, nil
}

// HasAccess implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) GetAccessLevel(ctx context.Context, channelId *uuid.UUID, characterId string) (chat.ChannelPermissionLevel, error) {
	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return chat.PermissionNone, err
	}

	rows, err := tx.Query(
		ctx,
		`SELECT chat_banned_until
		FROM chat_channels
		JOIN chat_channel_permissions 
			ON chat_channels.id = chat_channel_permissions.chat_channel_id
		WHERE (
			chat_channel_permissions.character_id = $1
			AND chat_channels.deleted_at IS NULL
		)
		LIMIT 1`,
		characterId,
	)
	if err != nil {
		return chat.PermissionNone, err
	}

	if rows.Next() {
		perm := chat.ChannelPermission{}
		err = rows.Scan(&perm.ChatBannedUntil)
		if err != nil {
			return chat.PermissionNone, err
		}
		return perm.Level(), nil
	}

	return chat.PermissionNone, nil
}

// SaveForCharacter implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	tagCharacter(ctx, characterId)
	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "DELETE FROM chat_channel_permissions WHERE character_id = $1", characterId)
	if err != nil {
		return err
	}

	err = r.addPermissionsForUser(ctx, tx, channelIds, characterId)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// BanCharacter implements ChatChannelPermissionRepository.
func (r *ccpPgxRepo) BanCharacter(ctx context.Context, characterId string, channelId *uuid.UUID, until *time.Time) error {
	tagCharacter(ctx, characterId)
	tx, err := r.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		"UPDATE chat_channel_permissions SET chat_banned_until = $1 WHERE character_id = $2 AND chat_channel_id = $3",
		until,
		characterId,
		channelId,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *ccpPgxRepo) addPermissionsForUser(ctx context.Context, tx pgx.Tx, channelIds []*uuid.UUID, characterId string) error {
	for _, channelId := range channelIds {
		_, err := tx.Exec(ctx, "INSERT INTO chat_channel_permissions (chat_channel_id, character_id) VALUES ($1, $2)", channelId, characterId)
		if err != nil {
			return err
		}
	}
	return nil
}
