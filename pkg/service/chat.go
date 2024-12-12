package service

import (
	"context"
	"errors"
	"sync"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type ChatService interface {
	ReceiveChannelMessages(ctx context.Context, channelId *uuid.UUID, receiverCharacterId string) (chan *chat.Message, error)
	ReceiveDirectMessage(ctx context.Context, targetCharacterId, receiverCharacterId string) (chan *chat.Message, error)

	SendChannelMessage(ctx context.Context, channelId *uuid.UUID, msg *chat.Message) error
	SendDirectMessage(ctx context.Context, targetCharacterId string, msg *chat.Message) error

	Shutdown(ctx context.Context) error
}

type chatService struct {
	kafkaBrokers       []string
	openChannelReaders map[uuid.UUID]map[string]*kafka.Reader
	openDirectReaders  map[string]map[string]*kafka.Reader

	shuttingDown bool

	mu sync.Mutex
	wg sync.WaitGroup
}

// Shutdown implements ChatService.
func (s *chatService) Shutdown(ctx context.Context) error {
	s.shuttingDown = true

	var errs error
	for _, readers := range s.openChannelReaders {
		for _, reader := range readers {
			err := reader.Close()
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}

	s.wg.Wait()
	s.shuttingDown = false
	return errs
}

func NewChatService(kafkaBrokers []string) ChatService {
	return &chatService{
		kafkaBrokers:       kafkaBrokers,
		openChannelReaders: make(map[uuid.UUID]map[string]*kafka.Reader),
		openDirectReaders:  make(map[string]map[string]*kafka.Reader),
	}
}

// ReceiveChannelMessages implements ChatService.
func (s *chatService) ReceiveChannelMessages(
	ctx context.Context,
	channelId *uuid.UUID,
	receiverCharacterId string,
) (chan *chat.Message, error) {
	if s.shuttingDown {
		return nil, errors.New("service is shutting down")
	}

	outChan := make(chan *chat.Message)
	s.wg.Add(1)
	go func() {
		defer close(outChan)
		defer s.wg.Done()

		reader, ok := s.openChannelReaders[*channelId][receiverCharacterId]
		if !ok {
			reader = kafka.NewReader(kafka.ReaderConfig{
				Brokers: s.kafkaBrokers,
				GroupID: receiverCharacterId,
				Topic:   getTopicForChannel(channelId),
				Logger:  kafka.LoggerFunc(log.Logger.Tracef),
			})
			s.mu.Lock()
			if _, ok := s.openChannelReaders[*channelId]; !ok {
				s.openChannelReaders[*channelId] = make(map[string]*kafka.Reader)
			}
			s.openChannelReaders[*channelId][receiverCharacterId] = reader
			s.mu.Unlock()
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				kafkaMessage, err := reader.ReadMessage(ctx)
				if err != nil {
					log.Logger.Errorf(
						"error reading message from kafka for channel %s and character %s: %v",
						channelId.String(),
						receiverCharacterId,
						err,
					)
					return
				}

				outChan <- chat.MessageFromKafkaMessage(&kafkaMessage)
			}
		}
	}()

	return outChan, nil
}

// ReceiveDirectMessage implements ChatService.
func (s *chatService) ReceiveDirectMessage(ctx context.Context, targetCharacterId, receiverUserId string) (chan *chat.Message, error) {
	if s.shuttingDown {
		return nil, errors.New("service is shutting down")
	}

	outChan := make(chan *chat.Message)
	s.wg.Add(1)
	go func() {
		defer close(outChan)
		defer s.wg.Done()

		reader, ok := s.openDirectReaders[targetCharacterId][receiverUserId]
		if !ok {
			reader = kafka.NewReader(kafka.ReaderConfig{
				Brokers: s.kafkaBrokers,
				GroupID: receiverUserId,
				Topic:   getTopicForDirect(targetCharacterId),
				Logger:  kafka.LoggerFunc(log.Logger.Tracef),
			})
			s.mu.Lock()
			if _, ok := s.openDirectReaders[targetCharacterId]; !ok {
				s.openDirectReaders[targetCharacterId] = make(map[string]*kafka.Reader)
			}
			s.openDirectReaders[targetCharacterId][receiverUserId] = reader
			s.mu.Unlock()
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				kafkaMessage, err := reader.ReadMessage(ctx)
				if err != nil {
					log.Logger.Errorf(
						"error reading message from kafka for direct message to %s for user %s: %v",
						targetCharacterId,
						receiverUserId,
						err,
					)
					return
				}

				outChan <- chat.MessageFromKafkaMessage(&kafkaMessage)
			}
		}
	}()

	return outChan, nil
}

// SendChannelMessage implements ChatService.
func (s *chatService) SendChannelMessage(ctx context.Context, channelId *uuid.UUID, msg *chat.Message) error {
	if s.shuttingDown {
		return errors.New("service is shutting down")
	}
	s.wg.Add(1)
	defer s.wg.Done()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    getTopicForChannel(channelId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})
	defer writer.Close()

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

// SendDirectMessage implements ChatService.
func (s *chatService) SendDirectMessage(ctx context.Context, targetCharacterId string, msg *chat.Message) error {
	if s.shuttingDown {
		return errors.New("service is shutting down")
	}
	s.wg.Add(1)
	defer s.wg.Done()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  s.kafkaBrokers,
		Topic:    getTopicForDirect(targetCharacterId),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
		Logger:   kafka.LoggerFunc(log.Logger.Tracef),
	})
	defer writer.Close()

	return writer.WriteMessages(ctx, *msg.ToKafkaMessage())
}

func getTopicForChannel(channelId *uuid.UUID) string {
	return "channel_" + channelId.String()
}

func getTopicForDirect(characterId string) string {
	return "direct_" + characterId
}
