package usecase

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
)

// Create function implement user use case.
func (uc useCase) Create(ctx context.Context, request dto.CreateUser) error {
	return uc.userRepository.Create(ctx, request.ToUserModel())
}
