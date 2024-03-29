package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	goKitEcho "github.com/paulusrobin/gogen-golib/go-kit/echo"
	httpServer "github.com/paulusrobin/gogen-golib/server/http"
	goPlaygroundV10 "github.com/paulusrobin/gogen-golib/validator/integrations/go-playground-v10"
	"github.com/paulusrobin/gogen/internal/config"
	greetingEncoding "github.com/paulusrobin/gogen/internal/pkg/greeting/encoding"
	greetingEndpoint "github.com/paulusrobin/gogen/internal/pkg/greeting/endpoint"
	userEncoding "github.com/paulusrobin/gogen/internal/pkg/user/encoding"
	userEndpoint "github.com/paulusrobin/gogen/internal/pkg/user/endpoint"
	"github.com/paulusrobin/gogen/pkg/postgres"
	"github.com/paulusrobin/gogen/pkg/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

type IServer interface {
	Run() error
	Shutdown() error
}

func initializeServer(sig chan os.Signal, cfg config.Config) (IServer, error) {
	validation := goPlaygroundV10.NewValidation()
	database, err := postgres.NewDatabase(postgres.Config{
		DSN:                   cfg.Postgres.DSN,
		LogLevel:              cfg.Postgres.LogLevel,
		MaxOpenConnections:    cfg.Postgres.MaxOpenConnections,
		MaxIdleConnections:    cfg.Postgres.MaxIdleConnections,
		MaxConnectionLifetime: cfg.Postgres.MaxConnectionLifetime,
	})
	if err != nil {
		return nil, err
	}

	redisCache, err := redis.NewCache(redis.Config{
		Address:  cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return nil, err
	}

	return httpServer.Server(sig,
		httpServer.Config{
			Port:             cfg.HTTP.Port,
			GracefulDuration: cfg.HTTP.GracefulDuration,
		},
		httpServer.RegisterRoute(func(ec *echo.Echo) {
			greeting := greetingEndpoint.NewEndpoint(cfg, validation)
			greet := ec.Group("/greetings")
			greet.GET("", goKitEcho.Handler(
				greeting.Greet(),
				goKitEcho.WithDecoder(greetingEncoding.DecodeGreetingRequest(validation)),
			))
		}),
		httpServer.RegisterRoute(func(ec *echo.Echo) {
			users := userEndpoint.NewEndpoint(cfg, validation, database, redisCache)
			user := ec.Group("/api/v1/users")
			user.POST("", goKitEcho.Handler(
				users.Create(),
				goKitEcho.WithDecoder(userEncoding.DecodeCreateRequest(validation)),
			))
			user.GET("/:id", goKitEcho.Handler(
				users.GetByID(),
				goKitEcho.WithDecoder(userEncoding.DecodeGetRequest(validation)),
			))
		}),
	)
}

func runner(cfg config.Config) func(c *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		if !cfg.IsValidHTTP() {
			return fmt.Errorf("invalid required config for http")
		}

		log.Info().Msgf("[http-server] starting server with [%s] log level", zerolog.GlobalLevel().String())

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		s, err := initializeServer(sigChannel, cfg)
		if err != nil {
			return err
		}

		go s.Run()

		sig := <-sigChannel
		log.Info().Msgf("[http-server] signal %s received, exiting", sig.String())
		if err = s.Shutdown(); err != nil {
			return err
		}
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
