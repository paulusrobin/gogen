package usecase

import (
	"github.com/paulusrobin/gogen/internal/pkg/greeting"
)

type useCase struct{}

// NewUseCase function to initialize user use case.
func NewUseCase() greeting.UseCase {
	return useCase{}
}
