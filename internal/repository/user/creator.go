package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/user/dto"
)

type Creator interface {
	Create(ctx context.Context, request dto.CreateRequest) (dto.CreateResponse, error)
}
