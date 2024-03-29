package usecase

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/dto"
)

// Greet function to implement use case.
func (u useCase) Greet(ctx context.Context, request dto.GreetingRequest) (dto.GreetingResponse, error) {
	var greeting = "welcome to the server"
	if request.Name != "" {
		greeting = "hi " + request.Name + ", " + greeting
	}
	return dto.GreetingResponse{
		Message: greeting,
	}, nil
}
