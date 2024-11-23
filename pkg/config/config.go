package config

import (
	"context"

	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
)

var (
	Version = "v1.0.0"
)

type ChatConfig struct {
	cconfig.BaseConfig `yaml:",inline" chatstructure:",squash"`
	Postgres           cconfig.DBPoolConfig `yaml:"postgres"`
}

func NewChatConfig(ctx context.Context) (*ChatConfig, error) {
	config := &ChatConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8085",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "localhost:8080",
				Realm:        "default",
				Id:           "7b575e9b-c687-4cdc-b210-67c59b5f380f",
				ClientId:     "sro-chat-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            0,
			OpenTelemtryAddress: "localhost:4317",
		},
		Postgres: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{},
				Name:          "chat-service",
				Username:      "postgres",
				Password:      "password",
			},
		},
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-chat", config)
	return config, err
}
