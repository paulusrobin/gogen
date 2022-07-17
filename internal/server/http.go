package server

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	goPlaygroundV10 "github.com/paulusrobin/gogen-golib/validator/integrations/go-playground-v10"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/user"
	userEndpoint "github.com/paulusrobin/gogen/internal/pkg/user/endpoint"
	userUseCase "github.com/paulusrobin/gogen/internal/pkg/user/usecase"
	"github.com/paulusrobin/gogen/internal/repository/postgres"
	userRepository "github.com/paulusrobin/gogen/internal/repository/postgres/user"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"os"
	"syscall"
	"time"
)

const defaultGracefulDuration = 12 * time.Second

type (
	httpServer struct {
		sig chan os.Signal
		cfg config.Config

		// dependencies
		ec         *echo.Echo
		db         *gorm.DB
		validation validator.Validation
		user       userPackage
	}
	userPackage struct {
		repository     postgres.UserRepository
		useCase        user.UseCase
		createEndpoint endpoint.Endpoint
	}
)

// init function to initialize dependencies.
func (s *httpServer) init() error {
	// initialize database
	db, err := postgres.NewDatabase(s.cfg.Postgres)
	if err != nil {
		return err
	}

	// initialize validator
	s.validation = goPlaygroundV10.NewValidation()

	// initialize repository
	s.user.repository = userRepository.NewRepository(db)

	// initialize use case
	s.user.useCase = userUseCase.NewUseCase(s.user.repository)

	// initialize endpoint
	s.user.createEndpoint = userEndpoint.CreateUserEndpoint(userEndpoint.CreateUserParam{
		Cfg:        s.cfg,
		UseCase:    s.user.useCase,
		Validation: s.validation,
	})

	return nil
}

// middlewares function to register middleware to http server.
func (s *httpServer) middlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.Recover(),
		middleware.Gzip(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		}),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: uuid.New().String}),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				log.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Msg("[http-server] request")
				return nil
			},
		}),
	}
}

// routes function to initialize http routes.
func (s *httpServer) routes() {

}

// errorHandler function to handle http error.
func (s *httpServer) errorHandler(err error, c echo.Context) {
	_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
}

// Run function to run http server.
func (s *httpServer) Run() error {
	var parameters = map[string]interface{}{"server": s}
	defer func() {
		log.Info().
			Fields(parameters).
			Msg("[http-server] terminating server")
		s.sig <- syscall.SIGTERM
	}()

	// initialize http
	log.Info().Fields(parameters).Msg("[http-server] server initialized")
	s.ec = echo.New()
	s.ec.Use(s.middlewares()...)
	s.ec.HTTPErrorHandler = s.errorHandler

	// register routes
	log.Info().Fields(parameters).Msg("[http-server] registering server routes")
	s.routes()

	// run http
	log.Info().Fields(parameters).Msg("[http-server] starting server")
	if err := s.ec.Start(fmt.Sprintf(":%d", s.cfg.HTTP.Port)); err != nil {
		log.Error().Err(err).Fields(parameters).Msg("[http-server] failed to start server")
		return err
	}
	return nil
}

// Shutdown function to close http server.
func (s *httpServer) Shutdown() error {
	if s.cfg.HTTP.GracefulDuration == 0 {
		s.cfg.HTTP.GracefulDuration = defaultGracefulDuration
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.HTTP.GracefulDuration)
	defer cancel()

	if err := s.ec.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

// HTTP functions to initialize http server.
func HTTP(sig chan os.Signal, cfg config.Config) (Server, error) {
	s := &httpServer{sig: sig, cfg: cfg}
	if err := s.init(); err != nil {
		log.Error().Err(err).
			Fields(map[string]interface{}{"config": cfg}).
			Msg("[http-server] failed to initialize server")
		return nil, err
	}
	return s, nil
}
