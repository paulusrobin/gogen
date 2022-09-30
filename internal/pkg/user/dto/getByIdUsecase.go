package dto

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type (
	GetByIDUsecaseRequest struct {
		Mandatory mandatory.Request
		UserID    uint
	}
	GetByIDUsecaseResponse struct {
		Mandatory mandatory.Request
		User      model.User
	}
)

func (response GetByIDUsecaseResponse) ToGetByIDEndpointResponse() GetByIDEndpointResponse {
	return GetByIDEndpointResponse{
		User: response.User,
	}
}
