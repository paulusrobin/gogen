package greeting

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/payload"
)

// UseCase interface of user package.
type UseCase interface {
	Greet(ctx context.Context, request payload.GreetingRequest) (payload.GreetingResponse, error)
}
