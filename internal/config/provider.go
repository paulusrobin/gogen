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
	if err := readFromConsul(cfg); err != nil {
		panic(err)
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

func readFromConsul(cfg Config) error {
	hotReloadViper := viper.New()
	if err := hotReloadViper.AddRemoteProvider(remoteConfigProvider, cfg.endpoint(), cfg.Consul.Key); err == nil {
		log.Fatal().Err(err).Msg("cannot add remote config from consul")
		return err
	}

	hotReloadViper.SetConfigType("json")
	if err := hotReloadViper.ReadRemoteConfig(); err != nil {
		log.Error().Err(err).Msg("cannot read remote business config from consul, fallback to local default config")
	} else {
		if err := hotReloadViper.Unmarshal(&runtimeBusinessConfig); err != nil {
			log.Error().Err(err).Msg("cannot parse business config from consul, fallback to local default config")
		} else {
			log.Info().Msg("successfully initialized business config from consul")
		}
	}

	log.Info().Msg("remote config was read successfully from consul. starting periodical updates...")

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			if err := hotReloadViper.WatchRemoteConfig(); err != nil {
				log.Error().Err(err).Msg("cannot read remote business config from consul, using last fetched config")
				continue
			}

			if err := hotReloadViper.Unmarshal(&runtimeBusinessConfig); err != nil {
				log.Error().Err(err).Msg("cannot parse business config from consul, using last fetched config")
			}

			time.Sleep(time.Second * 5)
		}
	}()
	return nil
}
