package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/greeting"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/payload"
)

type GreetingParam struct {
	Cfg        config.Config
	UseCase    greeting.UseCase
	Validation validator.Validation
}

// GreetingEndpoint endpoint.
func GreetingEndpoint(param GreetingParam) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		greetingRequest, ok := request.(payload.GreetingRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}

		response, err := param.UseCase.Greet(ctx, greetingRequest)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}