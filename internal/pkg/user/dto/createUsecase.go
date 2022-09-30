package dto

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type (
	CreateUsecaseRequest struct {
		Mandatory  mandatory.Request
		FirstName  string
		MiddleName string
		LastName   string
	}
	CreateUsecaseResponse struct {
		Mandatory mandatory.Request
		User      model.User
	}
)

func (response CreateUsecaseResponse) ToCreateEndpointResponse() CreateEndpointResponse {
	return CreateEndpointResponse{
		User: response.User,
	}
}
