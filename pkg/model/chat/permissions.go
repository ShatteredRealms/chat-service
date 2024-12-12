package chat

import (
	"time"

	"github.com/google/uuid"
)

type ChannelPermissionLevel int8

const (
	PermissionNone ChannelPermissionLevel = iota
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
	if cp.ChatBannedUntil == nil || cp.ChatBannedUntil.Before(time.Now()) {
		return PermissionReadSend
	}

	return PermissionRead
}
