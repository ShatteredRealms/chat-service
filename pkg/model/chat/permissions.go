package chat

import (
	"time"

	"github.com/ShatteredRealms/go-common-service/pkg/log"
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

var (
	PermBanTime = time.Unix(0, 0).UTC()
)

type ChannelPermission struct {
	ChannelId       *uuid.UUID `db:"channel_id" json:"channelId"`
	CharacterId     string     `db:"character_id" json:"characterId"`
	ChatBannedUntil *time.Time `db:"chat_banned_until" json:"chatBannedUntil"`
}
type ChannelPermissions []*ChannelPermission

func (cp *ChannelPermission) Level() ChannelPermissionLevel {
	log.Logger.Infof("Until: %v", cp.ChatBannedUntil)
	log.Logger.Infof("PermBanTime: %v", PermBanTime)

	if cp.ChatBannedUntil == nil {
		return PermissionReadSend
	}

	if cp.ChatBannedUntil.Equal(PermBanTime) {
		log.Logger.Info("Banned")
		return PermissionPermBan
	}

	if cp.ChatBannedUntil.After(time.Now().UTC()) {
		return PermissionRead
	}

	return PermissionReadSend
}
