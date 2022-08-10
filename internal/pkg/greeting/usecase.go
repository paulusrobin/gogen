package greeting

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/dto"
)

// UseCase interface of user package.
type UseCase interface {
	Greet(ctx context.Context, request dto.GreetingRequest) (dto.GreetingResponse, error)
}
