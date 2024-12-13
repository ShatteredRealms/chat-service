package chat

import (
	"time"

	"github.com/google/uuid"
)

type ChannelPermissionLevel int8

const (
	// Channel Permission Levels. The lower, the more restrictive.
	PermissionPermBan ChannelPermissionLevel = iota
	PermissionNone
	PermissionRead
	PermissionReadSend
)

type ChannelPermission struct {
	ChannelId       *uuid.UUID `db:"channel_id" json:"channelId"`
	CharacterId     string     `db:"character_id" json:"characterId"`
	ChatBannedUntil *time.Time `db:"chat_banned_until" json:"chatBannedUntil"`
}
type ChannelPermissions []*ChannelPermission

func (cp *ChannelPermission) Level() ChannelPermissionLevel {
	if cp.ChatBannedUntil == nil {
		return PermissionReadSend
	}

	if cp.ChatBannedUntil.Equal(time.Time{}) {
		return PermissionPermBan
	}

	if cp.ChatBannedUntil.After(time.Now().UTC()) {
		return PermissionRead
	}

	return PermissionReadSend
}
