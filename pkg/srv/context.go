package srv

import (
	"context"
	"fmt"

	"github.com/ShatteredRealms/chat-service/pkg/config"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
	"github.com/ShatteredRealms/chat-service/pkg/service"
	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/character/characterbus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/dimensionbus"
	commonrepo "github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/ShatteredRealms/go-common-service/pkg/srv"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
)

type ChatContext struct {
	*srv.Context

	ChatService                  service.ChatService
	ChatChannelService           service.ChatChannelService
	ChatChannelPermissionService service.ChatChannelPermissionService

	DimensionService dimensionbus.Service
	CharacterService characterbus.Service
}

func NewChatContext(ctx context.Context, cfg *config.ChatConfig, serviceName string) (*ChatContext, error) {
	chatCtx := &ChatContext{
		Context: commonsrv.NewContext(&cfg.BaseConfig, serviceName),
	}
	ctx, span := chatCtx.Tracer.Start(ctx, "context.chat.new")
	defer span.End()

	pg, err := commonrepo.ConnectDB(ctx, cfg.Postgres, cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	chatCtx.ChatService = service.NewChatService(
		cfg.Kafka.Addresses(),
	)
	chatCtx.ChatChannelService = service.NewChatChannelService(
		repository.NewChatChannelPostgresRepository(pg),
	)
	chatCtx.ChatChannelPermissionService = service.NewChatChannelPermissionService(
		repository.NewChatChannelPermissionPostgresRepository(pg),
	)

	chatCtx.DimensionService = dimensionbus.NewService(
		dimensionbus.NewPostgresRepository(pg),
		bus.NewKafkaMessageBusReader(cfg.Kafka, serviceName, dimensionbus.Message{}),
	)
	chatCtx.DimensionService.StartProcessing(ctx)

	chatCtx.CharacterService = characterbus.NewService(
		characterbus.NewPostgresRepository(pg),
		bus.NewKafkaMessageBusReader(cfg.Kafka, serviceName, characterbus.Message{}),
	)
	chatCtx.CharacterService.StartProcessing(ctx)

	return chatCtx, nil
}

func (c *ChatContext) Shutdown() {
	c.DimensionService.StopProcessing()
	c.CharacterService.StopProcessing()
}
