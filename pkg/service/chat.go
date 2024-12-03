package service

import (
	"context"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type ChatService interface {
	ReceiveChannelMessage(ctx context.Context, channelId *uuid.UUID, receiverCharacterId string) (*chat.Message, error)
	ReceiveDirectMessage(ctx context.Context, targetCharacterId, receiverCharacterId string) (*chat.Message, error)

	SendChannelMessage(ctx context.Context, channelId *uuid.UUID, msg *chat.Message) error
	SendDirectMessage(ctx context.Context, targetCharacterId string, msg *chat.Message) error
}

type chatService struct {
	kafkaBrokers []string
}

func NewChatService(kafkaBrokers []string) ChatService {
	return &chatService{
		kafkaBrokers: kafkaBrokers,
	}
}

// ReceiveChannelMessage implements ChatService.
func (s *chatService) ReceiveChannelMessage(ctx context.Context, channelId *uuid.UUID, receiverCharacterId string) (*chat.Message, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: s.kafkaBrokers,
		GroupID: receiverCharacterId,
		Topic:   getTopicForChannel(channelId),
		Logger:  kafka.LoggerFunc(log.Logger.Tracef),
	})

	kafkaMessage, err := reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return chat.MessageFromKafkaMessage(&kafkaMessage), nil
}

// ReceiveDirectMessage implements ChatService.
func (s *chatService) ReceiveDirectMessage(ctx context.Context, targetCharacterId, receiverUserId string) (*chat.Message, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: s.kafkaBrokers,
		GroupID: receiverUserId,
		Topic:   getTopicForDirect(targetCharacterId),
		Logger:  kafka.LoggerFunc(log.Logger.Tracef),
	})

	kafkaMessage, err := reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return chat.MessageFromKafkaMessage(&kafkaMessage), nil
}

// SendChannelMessage implements ChatService.
func (s *chatService) SendChannelMessage(ctx context.Context, channelId *uuid.UUID, msg *chat.Message) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    getTopicForChannel(channelId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

// SendDirectMessage implements ChatService.
func (s *chatService) SendDirectMessage(ctx context.Context, targetCharacterId string, msg *chat.Message) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    getTopicForDirect(targetCharacterId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

func getTopicForChannel(channelId *uuid.UUID) string {
	return "channel_" + channelId.String()
}

func getTopicForDirect(characterId string) string {
	return "direct_" + characterId
}
