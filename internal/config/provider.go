package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// InitConfig initializes configuration from .env file and returns config structure.
func InitConfig() Config {
	cfg, err := readConfigFromEnv()
	if err != nil {
		panic("cannot initialize .env config")
	}

	provideRuntimeConfig(cfg)
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
