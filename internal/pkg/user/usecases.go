package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
)

// UseCase interface of user package.
type UseCase interface {
	Create(ctx context.Context, request dto.CreateUser) error
}
