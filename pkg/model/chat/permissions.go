package chat

import "github.com/google/uuid"

type ChannelPermission struct {
	ChannelId   *uuid.UUID `gorm:"uniqueIndex:idx_permission" json:"channelId"`
	CharacterId string     `gorm:"uniqueIndex:idx_permission" json:"characterId"`
}
type ChannelPermissions []*ChannelPermission
