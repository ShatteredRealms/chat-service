package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ShatteredRealms/chat-service/pkg/config"
	"github.com/ShatteredRealms/chat-service/pkg/pb"
	"github.com/ShatteredRealms/chat-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/character/characterbus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/dimensionbus"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/telemetry"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Load configuration and setup server context
	cfg, err := config.NewChatConfig(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("loading config: %v", err)
		return
	}

	srvCtx, err := srv.NewChatContext(ctx, cfg, config.ServiceName)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("chat server context: %v", err)
		return
	}
	defer srvCtx.Shutdown()
	ctx, span := srvCtx.Tracer.Start(ctx, "main")
	defer span.End()

	log.Logger.WithContext(ctx).Infof("Starting %s service", config.ServiceName)

	// OpenTelemetry setup
	otelShutdown, err := telemetry.SetupOTelSDK(ctx, "chat", config.Version, cfg.OpenTelemtryAddress)
	defer func() {
		log.Logger.Infof("Shutting down")
		err = otelShutdown(context.Background())
		if err != nil {
			log.Logger.Warnf("Error shutting down: %v", err)
		}
	}()

	if err != nil {
		log.Logger.WithContext(ctx).Errorf("connecting to otel: %v", err)
		return
	}

	// Common gRPC server setup
	keycloakClient := gocloak.NewClient(cfg.Keycloak.BaseURL)
	grpcServer, gwmux := util.InitServerDefaults(keycloakClient, cfg.Keycloak.Realm)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Health Service
	commonpb.RegisterHealthServiceServer(grpcServer, commonsrv.NewHealthServiceServer())
	err = commonpb.RegisterHealthServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register health service handler endpoint: %v", err)
		return
	}

	// Bus service
	busService, err := commonsrv.NewBusServiceServer(
		ctx,
		*srvCtx.Context,
		map[bus.BusMessageType]bus.Resettable{
			characterbus.Message{}.GetType(): srvCtx.CharacterService.GetResetter(),
			dimensionbus.Message{}.GetType(): srvCtx.DimensionService.GetResetter(),
		},
		map[bus.BusMessageType]commonsrv.WriterResetCallback{},
	)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("create bus service: %v", err)
		return
	}

	commonpb.RegisterBusServiceServer(grpcServer, busService)
	err = commonpb.RegisterBusServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register bus service handler endpoint: %v", err)
		return
	}

	// Chat Service
	chatService, err := srv.NewChatServiceServer(ctx, srvCtx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("creating chat service: %v", err)
		return
	}
	pb.RegisterChatServiceServer(grpcServer, chatService)
	err = pb.RegisterChatServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register chat service handler endpoint: %v", err)
		return
	}

	// Setup Complete
	log.Logger.WithContext(ctx).Info("Initialization complete")
	span.End()
	srv, srvErr := util.StartServer(ctx, grpcServer, gwmux, cfg.Server.Address())

	select {
	case err := <-srvErr:
		if err != nil {
			log.Logger.Error(err)
		}

	case <-ctx.Done():
		log.Logger.Info("Server canceled by user input.")
		srv.Shutdown(ctx)
		stop()
	}

	log.Logger.Info("Server stopped.")
}
