package main

import (
	"github.com/paulusrobin/gogen/cmd/http"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	cfg config.Config
	cmd = &cobra.Command{
		Use:   "gogen",
		Short: "gogen",
	}
)

func main() {
	cfg = config.InitConfig()
	globalLogLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(globalLogLevel)

	cmd.AddCommand(
		http.Cmd(cfg), // add http command
	)

	// execute command
	if err := cmd.Execute(); err != nil {
		panic("can't execute cmd")
	}
}
