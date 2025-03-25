package chat

import (
	"errors"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/google/uuid"
)

// Channel represents a chat channel. If the dimension ID is not empty, the channel is a dimension channel.
// Otherwise, the channel is a global channel for all dimensions.
type Channel struct {
	Id        uuid.UUID  `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Name      string     `db:"name" json:"name"`

	// DimensionId is the dimension ID for the channel.
	// If the dimension ID is not empty, the channel is a dimension channel.
	// Otherwise, the dimension ID is empty, and the channel is a global channel for all dimensions.
	DimensionId *uuid.UUID `db:"dimension_id" json:"dimensionId"`

	// Public is true if the channel does not require a permission to join.
	Public bool `db:"public" json:"public"`
}
type Channels []*Channel

func (c *Channel) ToPb() *pb.ChatChannel {
	log.Logger.Infof("Updating chat channel with request: %+v", *c)
	out := &pb.ChatChannel{
		Id:     c.Id.String(),
		Name:   c.Name,
		Public: c.Public,
	}
	if c.DimensionId != nil {
		out.DimensionId = c.DimensionId.String()
	}

	return out
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
