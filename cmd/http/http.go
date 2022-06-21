package http

import (
	"fmt"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func runner(cfg config.Config) func(c *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		if !cfg.IsValidHTTP() {
			return fmt.Errorf("invalid required config for http")
		}

		log.Info().Msgf("[http-server] starting server with [%s] log level", zerolog.GlobalLevel().String())

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		go server.HTTP(sigChannel, cfg)

		sig := <-sigChannel
		log.Info().Msgf("[http-server] signal %s received, exiting", sig)
		return nil
	}
}

// Cmd expose command runner
func Cmd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Run http server",
		RunE:  runner(cfg),
	}
}
