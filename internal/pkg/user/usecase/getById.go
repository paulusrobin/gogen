package usecase

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
	userRepoDto "github.com/paulusrobin/gogen/internal/repository/user/dto"
)

func (u useCase) GetByID(ctx context.Context, request dto.GetByIDUsecaseRequest) (dto.GetByIDUsecaseResponse, error) {
	response, err := u.userRepo.GetByID(ctx, userRepoDto.GetByIDRequest{UserID: request.UserID})
	if err != nil {
		return dto.GetByIDUsecaseResponse{}, err
	}

	return dto.GetByIDUsecaseResponse{
		Mandatory: request.Mandatory,
		User:      *response.User,
	}, err
}
