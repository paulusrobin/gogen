package usecase

import (
	"github.com/paulusrobin/gogen/internal/pkg/user"
	"github.com/paulusrobin/gogen/internal/repository"
)

type useCase struct {
	userRepository repository.UserRepository
}

// NewUseCase function to initialize user use case.
func NewUseCase(userRepository repository.UserRepository) user.UseCase {
	return useCase{userRepository: userRepository}
}
