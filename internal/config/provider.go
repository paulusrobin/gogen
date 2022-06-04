package config

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // needed to use viper remote config features
)

const remoteConfigProvider = "consul"

// InitConfig initializes configuration from .env file and returns config structure.
func InitConfig() Config {
	cfg, err := readConfigFromEnv()
	if err != nil {
		panic("cannot initialize .env config")
	}
	readFromConsul(cfg)
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

func readFromConsul(cfg Config) {
	hotReloadViper := viper.New()
	if err := hotReloadViper.AddRemoteProvider(remoteConfigProvider, cfg.endpoint(), cfg.Consul.Key); err == nil {
		log.Error().
			Err(err).
			Msg("cannot read remote conf from consul, fallback to local default config")
		return
	}

	hotReloadViper.SetConfigType("json")
	if err := hotReloadViper.ReadRemoteConfig(); err != nil {
		log.Error().Err(err).Msg("cannot read remote conf from consul, fallback to local default config")
	} else {
		if err := hotReloadViper.Unmarshal(&RuntimeConfig); err != nil {
			log.Error().Err(err).Msg("cannot parse conf from consul, fallback to local default config")
		}
	}

	log.Info().Msg("remote config was read successfully from consul. starting periodical updates...")

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			if err := hotReloadViper.WatchRemoteConfig(); err != nil {
				log.Error().Err(err).Msg("cannot read remote conf from consul")
				continue
			}

			if err := hotReloadViper.Unmarshal(&RuntimeConfig); err != nil {
				log.Error().Err(err).Msg("cannot parse conf from consul, using last fetched config")
			}

			time.Sleep(time.Second * 5)
		}
	}()
}
