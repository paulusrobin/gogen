package dto

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type (
	CreateEndpointRequest struct {
		Mandatory mandatory.Request
	}
	CreateEndpointResponse struct {
		User model.User
	}
)

func (request CreateEndpointRequest) ToCreateUsecaseRequest() CreateUsecaseRequest {
	return CreateUsecaseRequest{
		Mandatory: request.Mandatory,
	}
}
