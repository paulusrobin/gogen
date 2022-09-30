package endpoint

import (
	"github.com/go-redis/redis/v8"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/user"
	"github.com/paulusrobin/gogen/internal/pkg/user/usecase"
	"github.com/paulusrobin/gogen/internal/repository/user/cached"
	"github.com/paulusrobin/gogen/internal/repository/user/database"
	"gorm.io/gorm"
)

type Endpoint struct {
	cfg        config.Config
	useCase    user.UseCase
	validation validator.Validation
}

// NewEndpoint function to initialize greeting endpoint.
func NewEndpoint(
	cfg config.Config,
	validation validator.Validation,
	db *gorm.DB,
	cache *redis.Client,
) user.Endpoint {
	return Endpoint{
		cfg:        cfg,
		validation: validation,
		useCase:    usecase.NewUseCase(cached.UserRepository(cache, database.UserRepository(db))),
	}
}
