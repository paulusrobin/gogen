package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/config"
	"github.com/paulusrobin/gogen/internal/pkg/user"
	"github.com/paulusrobin/gogen/internal/pkg/user/payload"
)

type CreateUserParam struct {
	Cfg        config.Config
	UseCase    user.UseCase
	Validation validator.Validation
}

func CreateUserEndpoint(param CreateUserParam) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createUserRequest, ok := request.(payload.CreateUser)
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}

		if err := param.UseCase.Create(ctx, createUserRequest); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
