package config

import (
	"context"

	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/sirupsen/logrus"
)

var (
	Version     = "v1.0.0"
	ServiceName = "ChatService"
)

type ChatConfig struct {
	cconfig.BaseConfig `yaml:",inline" chatstructure:",squash"`
	Postgres           cconfig.DBPoolConfig `yaml:"postgres"`
	Redis              cconfig.DBPoolConfig `yaml:"redis"`
}

func NewChatConfig(ctx context.Context) (*ChatConfig, error) {
	config := &ChatConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8180",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "http://localhost:8080",
				Realm:        "default",
				Id:           "a677cb5c-f24d-4f2e-aedd-b2f5387078e9",
				ClientId:     "sro-chat-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            logrus.InfoLevel,
			OpenTelemtryAddress: "localhost:4317",
			Kafka: cconfig.ServerAddresses{
				{
					Host: "localhost",
					Port: "29092",
				},
			},
		},
		Postgres: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "5432",
				},
				Name:     "chat_service",
				Username: "postgres",
				Password: "password",
			},
		},
		Redis: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "7000",
				},
			},
		},
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-chat", config)
	return config, err
}
