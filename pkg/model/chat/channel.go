package chat

import (
	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/model"
)

type Channel struct {
	model.Model
	Name        string `gorm:"uniqueIndex:idx_channel" json:"name"`
	DimensionId string `gorm:"uniqueIndex:idx_channel" json:"dimensionId"`
}
type Channels []*Channel

func (c *Channel) ToPb() *pb.ChatChannel {
	return &pb.ChatChannel{
		Id:            c.Id.String(),
		Name:          c.Name,
		DimensionId: c.DimensionId,
	}
}

func (c Channels) ToPb() *pb.ChatChannels {
	resp := &pb.ChatChannels{Channels: make([]*pb.ChatChannel, len(c))}
	for idx, channel := range c {
		resp.Channels[idx] = channel.ToPb()
	}

	return resp
}
