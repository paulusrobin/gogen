package server

import (
	"github.com/go-kit/kit/endpoint"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/user"
	userEndpoint "github.com/paulusrobin/gogen/internal/pkg/user/endpoint"
	userUseCase "github.com/paulusrobin/gogen/internal/pkg/user/usecase"
	"github.com/paulusrobin/gogen/internal/repository/postgres"
	userRepository "github.com/paulusrobin/gogen/internal/repository/postgres/user"
	"gorm.io/gorm"
	"os"
)

type (
	httpServer struct {
		sig chan os.Signal
		cfg config.Config

		// dependencies
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
	s.validation = nil

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

// routes function to initialize http routes.
func (s *httpServer) routes() {

}

func (s *httpServer) Run() error {
	if err := s.init(); err != nil {
		return err
	}
	return nil
}

func HTTP(sig chan os.Signal, cfg config.Config) Server {
	return &httpServer{
		sig: sig,
		cfg: cfg,
	}
}
