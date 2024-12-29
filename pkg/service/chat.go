package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type ChatService interface {
	ReceiveChannelMessages(
		ctx context.Context,
		channelId *uuid.UUID,
		dimensionId *uuid.UUID,
		receiverCharacterId *uuid.UUID,
	) (chan *chat.Message, error)
	ReceiveDirectMessages(
		ctx context.Context,
		targetCharacterId, receiverCharacterId *uuid.UUID,
	) (chan *chat.Message, error)

	SendChannelMessage(
		ctx context.Context,
		channelId *uuid.UUID,
		dimensionId *uuid.UUID,
		msg *chat.Message,
	) error
	SendDirectMessage(ctx context.Context, targetCharacterId *uuid.UUID, msg *chat.Message) error

	Shutdown(ctx context.Context) error
}

type chatService struct {
	kafkaBrokers []string
	readers      map[uuid.UUID]map[uuid.UUID]*kafka.Reader

	shuttingDown bool

	mu sync.Mutex
	wg sync.WaitGroup
}

// Shutdown implements ChatService.
func (s *chatService) Shutdown(ctx context.Context) error {
	s.shuttingDown = true

	var errs error
	for _, readers := range s.readers {
		for _, reader := range readers {
			err := reader.Close()
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}

	s.wg.Wait()
	s.readers = make(map[uuid.UUID]map[uuid.UUID]*kafka.Reader)
	s.shuttingDown = false
	return errs
}

func NewChatService(kafkaBrokers []string) ChatService {
	return &chatService{
		kafkaBrokers: kafkaBrokers,
		readers:      make(map[uuid.UUID]map[uuid.UUID]*kafka.Reader),
	}
}

// ReceiveChannelMessages implements ChatService.
func (s *chatService) ReceiveChannelMessages(
	ctx context.Context,
	channelId *uuid.UUID,
	dimensionId *uuid.UUID,
	receiverCharacterId *uuid.UUID,
) (chan *chat.Message, error) {
	if s.shuttingDown {
		return nil, errors.New("service is shutting down")
	}

	outChan := make(chan *chat.Message)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		reader, cleanup := s.getReader(channelId, receiverCharacterId, s.getTopicForChannel(channelId, dimensionId))
		defer cleanup()

		s.messageLoop(ctx, channelId, receiverCharacterId, reader, outChan)
	}()

	return outChan, nil
}

// ReceiveDirectMessage implements ChatService.
func (s *chatService) ReceiveDirectMessages(ctx context.Context, targetCharacterId, receiverUserId *uuid.UUID) (chan *chat.Message, error) {
	if s.shuttingDown {
		return nil, errors.New("service is shutting down")
	}

	outChan := make(chan *chat.Message)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		reader, cleanup := s.getReader(targetCharacterId, receiverUserId, s.getTopicForDirect(targetCharacterId))
		defer cleanup()

		s.messageLoop(ctx, targetCharacterId, receiverUserId, reader, outChan)
	}()

	return outChan, nil
}

func (s *chatService) messageLoop(
	ctx context.Context,
	channelId *uuid.UUID,
	receiverId *uuid.UUID,
	reader *kafka.Reader,
	outChan chan *chat.Message,
) {
	defer close(outChan)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			kafkaMessage, err := reader.ReadMessage(ctx)
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					log.Logger.Errorf(
						"error reading message from kafka for channel %s and character %s: %v",
						channelId.String(),
						receiverId.String(),
						err,
					)
				}
				return
			}

			outChan <- chat.MessageFromKafkaMessage(&kafkaMessage)
		}
	}
}

// SendChannelMessage implements ChatService.
func (s *chatService) SendChannelMessage(
	ctx context.Context,
	channelId *uuid.UUID,
	dimensionId *uuid.UUID,
	msg *chat.Message,
) error {
	if s.shuttingDown {
		return errors.New("service is shutting down")
	}
	s.wg.Add(1)
	defer s.wg.Done()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    s.getTopicForChannel(channelId, dimensionId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})
	defer writer.Close()

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

// SendDirectMessage implements ChatService.
func (s *chatService) SendDirectMessage(ctx context.Context, targetCharacterId *uuid.UUID, msg *chat.Message) error {
	if s.shuttingDown {
		return errors.New("service is shutting down")
	}
	s.wg.Add(1)
	defer s.wg.Done()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    s.getTopicForDirect(targetCharacterId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})
	defer writer.Close()

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

func (s *chatService) getReader(channelId, receiverId *uuid.UUID, topic string) (*kafka.Reader, func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	reader, ok := s.readers[*channelId][*receiverId]
	if !ok {
		reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers: s.kafkaBrokers,
			// GroupID: channelId.String(),
			Topic:       topic,
			Logger:      kafka.LoggerFunc(log.Logger.Tracef),
			StartOffset: kafka.LastOffset,
		})
		if _, ok := s.readers[*channelId]; !ok {
			s.readers[*channelId] = make(map[uuid.UUID]*kafka.Reader)
		}
		s.readers[*channelId][*receiverId] = reader
	}
	cleanup := func() {
		reader.Close()
		delete(s.readers[*channelId], *receiverId)
	}
	return reader, cleanup
}

func (s *chatService) getTopicForChannel(channelId *uuid.UUID, dimensionId *uuid.UUID) (topic string) {
	topic = fmt.Sprintf("chat.channel.%s.%s", channelId.String(), dimensionId.String())
	s.configureChatTopic(topic)
	return
}

func (s *chatService) getTopicForDirect(characterId *uuid.UUID) (topic string) {
	topic = fmt.Sprintf("chat.direct.%s", characterId.String())
	s.configureChatTopic(topic)
	return topic
}

func (s *chatService) configureChatTopic(name string) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		client := kafka.Client{
			Addr: kafka.TCP(s.kafkaBrokers[0]),
		}
		resp, err := client.AlterConfigs(context.Background(), &kafka.AlterConfigsRequest{
			Resources: []kafka.AlterConfigRequestResource{
				{
					ResourceType: kafka.ResourceTypeTopic,
					ResourceName: name,
					Configs: []kafka.AlterConfigRequestConfig{
						{
							Name:  "retention.ms",
							Value: "2000",
						},
						{
							Name:  "delete.retention.ms",
							Value: "1000",
						},
						{
							Name:  "segment.ms",
							Value: "4000",
						},
					},
				},
			},
			ValidateOnly: false,
		})

		if err != nil {
			log.Logger.Errorf("error updating topic %s: %v", name, err)
			return
		}

		for resource, err := range resp.Errors {

			if err != nil {
				log.Logger.Errorf("error updating topic %s: %s", resource.Name, err.Error())
			}
		}
	}()
}
