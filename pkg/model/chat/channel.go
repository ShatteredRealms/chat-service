package chat

import (
	"errors"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/google/uuid"
)

// Channel represents a chat channel. If the dimension ID is not empty, the channel is a dimension channel.
// Otherwise, the channel is a global channel for all dimensions.
type Channel struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid()" db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `gorm:"uniqueIndex:idx_deleted" db:"deleted_at"`
	Name      string     `gorm:"uniqueIndex:idx_channel" json:"name" db:"name"`

	// DimensionId is the dimension ID for the channel.
	// If the dimension ID is not empty, the channel is a dimension channel.
	// Otherwise, the dimension ID is empty, and the channel is a global channel for all dimensions.
	DimensionId *uuid.UUID `gorm:"uniqueIndex:idx_channel" json:"dimensionId" db:"dimension_id"`
}
type Channels []*Channel

func (c *Channel) ToPb() *pb.ChatChannel {
	return &pb.ChatChannel{
		Id:          c.Id.String(),
		Name:        c.Name,
		DimensionId: c.DimensionId.String(),
	}
}

func (c Channels) ToPb() *pb.ChatChannels {
	resp := &pb.ChatChannels{Channels: make([]*pb.ChatChannel, len(c))}
	for idx, channel := range c {
		resp.Channels[idx] = channel.ToPb()
	}

	return resp
}

func (c *Channel) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
