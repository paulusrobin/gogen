package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
)

// UseCase interface of user package.
type UseCase interface {
	Create(ctx context.Context, request dto.CreateUsecaseRequest) (dto.CreateUsecaseResponse, error)
	GetByID(ctx context.Context, request dto.GetByIDUsecaseRequest) (dto.GetByIDUsecaseResponse, error)
}
