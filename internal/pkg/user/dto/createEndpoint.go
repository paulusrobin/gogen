package dto

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type (
	CreateEndpointRequest struct {
		Mandatory  mandatory.Request
		FirstName  string `json:"first_name" validate:"required"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
	}
	CreateEndpointResponse struct {
		User model.User
	}
)

func (request CreateEndpointRequest) ToCreateUsecaseRequest() CreateUsecaseRequest {
	return CreateUsecaseRequest{
		Mandatory:  request.Mandatory,
		FirstName:  request.FirstName,
		MiddleName: request.MiddleName,
		LastName:   request.LastName,
	}
}
