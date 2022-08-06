package usecase

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/payload"
)

// Greet function to implement use case.
func (u useCase) Greet(ctx context.Context, request payload.GreetingRequest) (payload.GreetingResponse, error) {
	var greeting = "welcome to the server"
	if request.Name != "" {
		greeting = "hi " + request.Name + ", " + greeting
	}
	return payload.GreetingResponse{
		Message: greeting,
	}, nil
}
