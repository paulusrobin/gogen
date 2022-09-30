package usecase

import (
	"context"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
	"github.com/paulusrobin/gogen/internal/repository/model"
	userRepoDto "github.com/paulusrobin/gogen/internal/repository/user/dto"
)

func (u useCase) Create(ctx context.Context, request dto.CreateUsecaseRequest) (dto.CreateUsecaseResponse, error) {
	userData := model.User{
		FirstName:  request.FirstName,
		MiddleName: request.MiddleName,
		LastName:   request.LastName,
	}
	response, err := u.userRepo.Create(ctx, userRepoDto.CreateRequest{User: &userData})
	if err != nil {
		return dto.CreateUsecaseResponse{}, err
	}

	return dto.CreateUsecaseResponse{
		Mandatory: request.Mandatory,
		User:      *response.User,
	}, err
}
