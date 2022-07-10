package usecase

import (
	"github.com/paulusrobin/gogen/internal/pkg/user"
	"github.com/paulusrobin/gogen/internal/repository/postgres"
)

type useCase struct {
	userRepository postgres.UserRepository
}

// NewUseCase function to initialize user use case.
func NewUseCase(userRepository postgres.UserRepository) user.UseCase {
	return useCase{userRepository: userRepository}
}
