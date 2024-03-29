package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/dto"
)

func (e Endpoint) Greet() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		greetingRequest, ok := request.(dto.GreetingRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}

		response, err := e.useCase.Greet(ctx, greetingRequest)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
