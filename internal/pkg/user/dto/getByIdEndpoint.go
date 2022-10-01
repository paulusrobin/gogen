package dto

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type (
	GetByIDEndpointRequest struct {
		Mandatory mandatory.Request
		UserID    uint `param:"id"`
	}
	GetByIDEndpointResponse struct {
		User model.User
	}
)

func (request GetByIDEndpointRequest) ToGetByIDUsecaseRequest() GetByIDUsecaseRequest {
	return GetByIDUsecaseRequest{
		Mandatory: request.Mandatory,
		UserID:    request.UserID,
	}
}
