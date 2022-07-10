package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/payload"
)

// UseCase interface of user package.
type UseCase interface {
	Create(ctx context.Context, request payload.CreateUser) error
}
