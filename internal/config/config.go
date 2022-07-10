package config

import (
	"fmt"
	"os"
	"strconv"
)

type (
	// Config project config.
	Config struct {
		Consul               Consul `mapstructure:",squash"`
		HTTP                 HTTP   `mapstructure:",squash"`
		GRPC                 GRPC   `mapstructure:",squash"`
		PubSub               PubSub `mapstructure:",squash"`
		PProf                Pprof  `mapstructure:",squash"`
		RuntimeConfigEnabled bool   `mapstructure:"RUNTIME_CONFIG_ENABLED"`
	}

	// Consul holds consul remote config related values.
	Consul struct {
		Host  string `mapstructure:"CONSUL_HOST"`
		Port  int    `mapstructure:"CONSUL_PORT"`
		Key   string `mapstructure:"CONSUL_KEY"`
		Token string `mapstructure:"CONSUL_HTTP_TOKEN"`
	}

	// HTTP contains HTTP related configuration.
	HTTP struct {
		Port int `mapstructure:"HTTP_PORT"`
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
