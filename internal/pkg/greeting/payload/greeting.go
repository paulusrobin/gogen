package payload

import (
	"github.com/labstack/echo/v4"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
)

type (
	GreetingRequest struct {
		Name string `query:"name"`
	}
	GreetingResponse struct {
		Message string `json:"message"`
	}
)

// DecodeGreetingRequest decoder function for greeting endpoint.
func DecodeGreetingRequest(validator validator.Validation) func(c echo.Context) (interface{}, error) {
	return func(c echo.Context) (interface{}, error) {
		var request GreetingRequest
		if err := c.Bind(&request); err != nil {
			return nil, err
		}
		return request, nil
	}
}
