package endpoint

import (
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/greeting"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/usecase"
)

type Endpoint struct {
	cfg        config.Config
	useCase    greeting.UseCase
	validation validator.Validation
}

// NewEndpoint function to initialize greeting endpoint.
func NewEndpoint(cfg config.Config, validation validator.Validation) Endpoint {
	return Endpoint{
		cfg:        cfg,
		validation: validation,
		useCase:    usecase.NewUseCase(),
	}
}
