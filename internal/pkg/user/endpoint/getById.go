package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
)

func (e Endpoint) GetByID() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		payload, ok := request.(dto.GetByIDEndpointRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}
		return e.useCase.GetByID(ctx, payload.ToGetByIDUsecaseRequest())
	}
}
