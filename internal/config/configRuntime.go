package config

import (
	consulViper "github.com/paulusrobin/gogen-golib/remote-config/consul/integrations/viper"
	consul "github.com/paulusrobin/gogen-golib/remote-config/consul/interface"
	"github.com/pkg/errors"
	"sync"
)

type (
	businessConfig struct {
		sync.Mutex
	}
)

var runtimeBusinessConfig = newBusinessConfig()

func newBusinessConfig() businessConfig {
	// fallback values
	return businessConfig{}
}

func provideRuntimeConfig(cfg Config) {
	reader, err := consulViper.NewConsulReader(consul.Config{
		Connection: consul.ConnectionConfig{
			Host:  cfg.Consul.Host,
			Port:  cfg.Consul.Port,
			Token: cfg.Consul.Token,
			Key:   cfg.Consul.Key,
		},
		ConfigType: "json",
		Interval:   cfg.Consul.RefreshInterval,
	})
	if err != nil {
		panic(errors.Wrap(err, "cannot initialize consul reader"))
	}

	if err = reader.Read(&runtimeBusinessConfig); err != nil {
		panic(errors.Wrap(err, "cannot initialize consul business config variable"))
	}
}
