package payload

import (
	"github.com/labstack/echo/v4"
	"github.com/paulusrobin/gogen-golib/mandatory"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"net/http"
)

type CreateUser struct {
	Mandatory  mandatory.Request
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (cu CreateUser) ToUserModel() *model.User {
	return &model.User{
		FirstName:  cu.FirstName,
		MiddleName: cu.MiddleName,
		LastName:   cu.LastName,
	}
}

func DecodeCreateRequest(validator validator.Validation) func(c echo.Context) (interface{}, error) {
	return func(c echo.Context) (interface{}, error) {
		var request CreateUser
		if err := c.Bind(&request); err != nil {
			return nil, err
		}
		if err := validator.Validator(request.Mandatory.Language()).Struct(request); err != nil {
			return nil, err
		}
		return request, nil
	}
}

func EncodeCreateRequest(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": response,
	})
}
