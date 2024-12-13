package chat

import (
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/google/uuid"
)

type MessageLog struct {
	SentAt      time.Time  `db:"sent_at" json:"sentAt"`
	ChannelId   *uuid.UUID `db:"chat_channel_id" json:"channelId"`
	CharacterId *uuid.UUID `db:"character_id" json:"characterId"`
	Content     string     `db:"content" json:"content"`
}

type MessageLogs []*MessageLog

func (m *MessageLog) ToPb() *pb.ChatLog {
	return &pb.ChatLog{
		SentAt:    uint64(m.SentAt.Unix()),
		ChannelId: m.ChannelId.String(),
		SenderId:  m.CharacterId.String(),
		Content:   m.Content,
	}
}

func (m MessageLogs) ToPb() *pb.ChatLogs {
	resp := &pb.ChatLogs{Logs: make([]*pb.ChatLog, len(m))}
	for idx, log := range m {
		resp.Logs[idx] = log.ToPb()
	}

	return resp
}
