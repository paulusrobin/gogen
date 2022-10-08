package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/user/dto"
)

type GetterByID interface {
	GetByID(ctx context.Context, request dto.GetByIDRequest) (dto.GetByIDResponse, error)
}
