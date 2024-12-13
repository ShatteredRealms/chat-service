package chat

import "github.com/segmentio/kafka-go"

type Message struct {
	SenderCharacterId string `json:"senderCharacterId"`
	Content           string `json:"content"`
}
type Messages []*Message

func (m *Message) ToKafkaMessage() *kafka.Message {
	return &kafka.Message{
		Key:   []byte(m.SenderCharacterId),
		Value: []byte(m.Content),
	}
}

func MessageFromKafkaMessage(msg *kafka.Message) *Message {
	return &Message{
		SenderCharacterId: string(msg.Key),
		Content:           string(msg.Value),
	}
}
