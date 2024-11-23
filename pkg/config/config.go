package config

import "github.com/ShatteredRealms/go-common-service/pkg/config"

var (
	Version = "v1.0.0"
)

type ChatConfig struct {
	config.BaseConfig `yaml:",inline" mapstructure:",squash"`
	Postgres          config.DBPoolConfig `yaml:"postgres"`
}

func NewChatConfig() *ChatConfig {
	return &ChatConfig{
		BaseConfig: config.BaseConfig{
			Server: config.ServerAddress{
				Host: "localhost",
				Port: "8180",
			},
			Keycloak: config.KeycloakConfig{
				BaseURL:      "localhost:8080",
				Realm:        "default",
				Id:           "780b129c-3f75-441c-87ab-ace6c5691bd8",
				ClientId:     "sro-chat-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            0,
			OpenTelemtryAddress: "localhost:4317",
		},
		Postgres: config.DBPoolConfig{
			Master: config.DBConfig{
				ServerAddress: config.ServerAddress{},
				Name:          "chat-service",
				Username:      "postgres",
				Password:      "password",
			},
		},
	}
}
