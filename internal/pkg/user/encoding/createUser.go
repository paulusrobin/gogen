package encoding

import (
	"github.com/labstack/echo/v4"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/pkg/user/dto"
)

// DecodeCreateRequest decoder function for create endpoint.
func DecodeCreateRequest(validator validator.Validation) func(c echo.Context) (interface{}, error) {
	return func(c echo.Context) (interface{}, error) {
		var request dto.CreateEndpointRequest
		if err := c.Bind(&request); err != nil {
			return nil, err
		}
		if err := validator.Validator(request.Mandatory.Language()).Struct(request); err != nil {
			return nil, err
		}
		return request, nil
	}
}
