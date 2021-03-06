package grpc

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
		if !cfg.IsValidGRPC() {
			return fmt.Errorf("invalid required config for grpc")
		}

		log.Info().Msgf("[grpc-server] starting server with [%s] log level", zerolog.GlobalLevel().String())

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		s, err := server.GRPC(sigChannel, cfg)
		if err != nil {
			return err
		}

		go s.Run()

		sig := <-sigChannel
		log.Info().Msgf("[grpc-server] signal %s received, exiting", sig)
		if err = s.Shutdown(); err != nil {
			return err
		}
		return nil
	}
}

// Cmd expose command runner
func Cmd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "Run grpc server",
		RunE:  runner(cfg),
	}
}
