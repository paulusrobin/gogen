package config

import (
	consulViper "github.com/paulusrobin/gogen-golib/remote-config/consul/integrations/viper"
	consul "github.com/paulusrobin/gogen-golib/remote-config/consul/interface"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// InitConfig initializes configuration from .env file and returns config structure.
func InitConfig() Config {
	var (
		cfg    Config
		reader consul.Reader
		err    error
	)

	cfg, err = readConfigFromEnv()
	if err != nil {
		panic("cannot initialize .env config")
	}

	reader, err = consulViper.NewConsulReader(consul.Config{
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

	return cfg
}

func readConfigFromEnv() (Config, error) {
	var cfg Config

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("error while reading .env config file")
		return Config{}, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to parsing .env config")
		return Config{}, err
	}
	return cfg, nil
}
