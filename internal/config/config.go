package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

type (
	// Config project config.
	Config struct {
		Consul   Consul           `mapstructure:",squash"`
		HTTP     HTTP             `mapstructure:",squash"`
		GRPC     GRPC             `mapstructure:",squash"`
		PubSub   PubSub           `mapstructure:",squash"`
		PProf    Pprof            `mapstructure:",squash"`
		Postgres PostgresDatabase `mapstructure:",squash"`
		Redis    Redis            `mapstructure:",squash"`
	}

	// PostgresDatabase contains postgres configuration.
	PostgresDatabase struct {
		DSN                   string          `mapstructure:"DB_POSTGRES_DSN"`
		LogLevel              logger.LogLevel `mapstructure:"DB_POSTGRES_LOG_LEVEL"`
		MaxOpenConnections    int             `mapstructure:"DB_POSTGRES_MAX_OPEN_CONNECTIONS"`
		MaxIdleConnections    int             `mapstructure:"DB_POSTGRES_MAX_IDLE_CONNECTIONS"`
		MaxConnectionLifetime time.Duration   `mapstructure:"DB_POSTGRES_MAX_CONNECTIONS_LIFETIME"`
	}

	// Redis contains redis configuration.
	Redis struct {
		Address  string `mapstructure:"REDIS_ADDRESS"`
		Password string `mapstructure:"REDIS_PASSWORD"`
		DB       int    `mapstructure:"REDIS_DB"`
	}

	// Consul contains consul remote config related values.
	Consul struct {
		Host  string `mapstructure:"CONSUL_HOST"`
		Port  int    `mapstructure:"CONSUL_PORT"`
		Token string `mapstructure:"CONSUL_HTTP_TOKEN"`

		// RuntimeConfig settings
		Key             string        `mapstructure:"CONSUL_KEY"`
		RefreshInterval time.Duration `mapstructure:"CONSUL_REFRESH_INTERVAL"`
	}

	// HTTP contains HTTP related configuration.
	HTTP struct {
		Port             int           `mapstructure:"HTTP_PORT"`
		GracefulDuration time.Duration `mapstructure:"HTTP_GRACEFUL_DURATION"`
	}

	// GRPC contains GRPC related configuration.
	GRPC struct {
		Port int `mapstructure:"GRPC_PORT"`
	}

	// PubSub contains google PubSub project configuration.
	PubSub struct {
		ProjectID   string `mapstructure:"PUBSUB_PROJECT_ID"`
		Credentials string `mapstructure:"PUBSUB_CREDENTIALS"`
	}

	// Pprof contains Pprof project configuration.
	Pprof struct {
		Enabled bool `mapstructure:"PPROF_ENABLED"`
		Port    int  `mapstructure:"PPROF_PORT"`
	}
)

func (c *Config) endpoint() string {
	if os.Getenv("CONSUL_HTTP_TOKEN") == "" && c.Consul.Token != "" {
		_ = os.Setenv("CONSUL_HTTP_TOKEN", c.Consul.Token)
	}

	if c.Consul.Host == "" {
		c.Consul.Host = os.Getenv("CONSUL_HOST")
	}

	if c.Consul.Port == 0 {
		consulPort := os.Getenv("CONSUL_PORT")
		if consulPort != "" {
			if port, err := strconv.Atoi(consulPort); err == nil {
				c.Consul.Port = port
			}
		}
	}

	return fmt.Sprintf("%s:%d", c.Consul.Host, c.Consul.Port)
}

func (c Config) IsValidHTTP() bool {
	return c.HTTP.Port != 0
}

func (c Config) IsValidGRPC() bool {
	return c.GRPC.Port != 0
}

func (c Config) IsValidSubscriber() bool {
	return c.PubSub.ProjectID != "" && c.PubSub.Credentials != ""
}
