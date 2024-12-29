package chat

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"
)

type messageMeta struct {
	TraceId    string `json:"traceId"`
	SentTimeMs int64  `json:"sentTimeMs"`
}

type messageData struct {
	SenderCharacterId string `json:"senderCharacterId"`
	Content           string `json:"content"`
}

type Message struct {
	messageMeta
	messageData
}
type Messages []*Message

func NewMessage(ctx context.Context, sender, content string) *Message {
	return &Message{
		messageMeta: messageMeta{
			TraceId:    trace.SpanFromContext(ctx).SpanContext().TraceID().String(),
			SentTimeMs: time.Now().UnixMilli(),
		},
		messageData: messageData{
			SenderCharacterId: sender,
			Content:           content,
		},
	}
}

func (m *Message) ToKafkaMessage() *kafka.Message {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(m.messageData)

	return &kafka.Message{
		Key:   []byte(m.TraceId),
		Value: buf.Bytes(),
		Time:  time.UnixMilli(m.SentTimeMs),
	}
}

func (m *Message) ToPb() *pb.ChatMessage {
	return &pb.ChatMessage{
		SenderCharacterId: m.SenderCharacterId,
		Content:           m.Content,
		SentTimeMs:        m.SentTimeMs,
	}
}

func MessageFromKafkaMessage(msg *kafka.Message) *Message {
	dec := gob.NewDecoder(bytes.NewBuffer(msg.Value))
	var chatMsg Message
	dec.Decode(&chatMsg)
	chatMsg.TraceId = string(msg.Key)
	chatMsg.SentTimeMs = msg.Time.UnixMilli()
	return &chatMsg
}
