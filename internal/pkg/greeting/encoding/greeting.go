package encoding

import (
	"github.com/labstack/echo/v4"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/pkg/greeting/dto"
)

// DecodeGreetingRequest decoder function for greeting endpoint.
func DecodeGreetingRequest(validator validator.Validation) func(c echo.Context) (interface{}, error) {
	return func(c echo.Context) (interface{}, error) {
		var request dto.GreetingRequest
		if err := c.Bind(&request); err != nil {
			return nil, err
		}
		return request, nil
	}
}
