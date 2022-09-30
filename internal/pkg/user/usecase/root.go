package usecase

import (
	"github.com/paulusrobin/gogen/internal/pkg/user"
	userRepository "github.com/paulusrobin/gogen/internal/repository/user"
)

type useCase struct {
	userRepo userRepository.Repository
}

// NewUseCase function to initialize user use case.
func NewUseCase(userRepo userRepository.Repository) user.UseCase {
	return useCase{userRepo: userRepo}
}
