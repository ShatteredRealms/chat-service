package repository

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type chatLogPgxRepository struct {
	conn *pgxpool.Pool
}

func NewChatLogPgxRepository(migrater *PgxMigrater) ChatLogRepository {
	return &chatLogPgxRepository{
		conn: migrater.conn,
	}
}

// AddMessage implements ChatLogRepository.
func (c *chatLogPgxRepository) AddMessage(ctx context.Context, channelId *uuid.UUID, dimensionId *uuid.UUID, message *chat.Message) error {
	if message == nil {
		return ErrEmptyMessage
	}

	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	senderId, err := uuid.Parse(message.SenderCharacterId)
	if err != nil {
		return err
	}
	ct, err := tx.Exec(ctx,
		"INSERT INTO chat_messages (channel_id, character_id, content) VALUES ($1, $2, $3)",
		channelId, &senderId, message.Content)
	if err != nil {
		return err
	}
	if ct.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return tx.Commit(ctx)
}

// GetMessages implements ChatLogRepository.
func (c *chatLogPgxRepository) GetMessages(
	ctx context.Context,
	channelId *uuid.UUID,
	limit *uint,
	offset *uint,
	before *time.Time,
	after *time.Time,
	sender *uuid.UUID,
) (outMsgs *chat.MessageLogs, count uint, err error) {
	// Validate input
	if after != nil && before != nil && after.Before(*before) {
		err = ErrInvalidTimeRange
		return
	}

	if limit == nil && offset != nil {
		err = ErrInvalidOffset
		return
	}

	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback(ctx)

	// Count messages to determine max offset
	whereBuilder := strings.Builder{}
	args := []any{channelId}
	getParamIndex := func() string {
		return strconv.Itoa(len(args))
	}
	whereBuilder.WriteString("WHERE (channel_id = $1")
	if after != nil {
		args = append(args, *after)
		whereBuilder.WriteString(" AND created_at > $")
		whereBuilder.WriteString(getParamIndex())
	}
	if before != nil {
		args = append(args, *before)
		whereBuilder.WriteString(" AND created_at < $")
		whereBuilder.WriteString(getParamIndex())
	}
	if sender != nil {
		args = append(args, *sender)
		whereBuilder.WriteString(" AND sender_id = $")
		whereBuilder.WriteString(getParamIndex())
	}
	whereBuilder.WriteString(")")

	countBuilder := strings.Builder{}
	countBuilder.WriteString("SELECT COUNT(*) FROM chat_messages ")
	countBuilder.WriteString(whereBuilder.String())

	err = tx.QueryRow(ctx, countBuilder.String(), args...).Scan(&count)
	if err != nil {
		return
	}

	maxOffset := uint(0)
	if limit != nil {
		// Formula: ((total + limit - 1) / limit) - 1
		// Example 1:
		//   20 messages
		//   Limit 10
		//   Offset 0: 1-10
		//   Offset 1: 11-20
		//   Max offset = 1
		//   uint((20 + 10 - 1) / 10) - 1
		//   == uint(29 / 10) - 1
		//   == uint(2.9) - 1
		//	 == 2 - 1
		//   == 1
		// Example 2:
		//	 21 messages
		//	 Limit 10
		//	 Offset 0: 1-10
		//	 Offset 1: 11-20
		//	 Offset 2: 21-30
		//	 Max offset = 2
		//	 uint((21 + 10 - 1) / 10) - 1
		//	 == uint(30 / 10) - 1
		//	 == uint(3) - 1
		//	 == 2
		maxOffset = ((uint(count) + *limit + 1) / *limit) - 1
	}

	if offset != nil && *offset > maxOffset {
		err = ErrInvalidOffset
		return
	}

	// Get messages
	getBuilder := strings.Builder{}
	getBuilder.WriteString("SELECT sent_at, channel_id, character_id, content FROM chat_messages ")
	getBuilder.WriteString(whereBuilder.String())
	getBuilder.WriteString(" ORDER BY sent_at DESC")
	if limit != nil && *limit > 0 {
		args = append(args, *limit)
		getBuilder.WriteString(" LIMIT $")
		getBuilder.WriteString(getParamIndex())
	}
	if offset != nil && *offset > 0 {
		args = append(args, *offset)
		getBuilder.WriteString(" OFFSET $")
		getBuilder.WriteString(getParamIndex())
	}

	rows, err := tx.Query(ctx, getBuilder.String(), args...)
	if err != nil {
		return
	}
	msgs := make(chat.MessageLogs, 0)
	for rows.Next() {
		msg := &chat.MessageLog{}
		err = rows.Scan(&msg.SentAt, &msg.ChannelId, &msg.CharacterId, &msg.Content)
		if err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return
	}

	outMsgs = &msgs
	return
}
