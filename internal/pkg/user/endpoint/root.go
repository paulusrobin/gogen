package endpoint

import (
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/user"
	"github.com/paulusrobin/gogen/internal/pkg/user/usecase"
)

type Endpoint struct {
	cfg        config.Config
	useCase    user.UseCase
	validation validator.Validation
}

// NewEndpoint function to initialize greeting endpoint.
func NewEndpoint(cfg config.Config, validation validator.Validation) user.Endpoint {
	return Endpoint{
		cfg:        cfg,
		validation: validation,
		useCase:    usecase.NewUseCase(),
	}
}
