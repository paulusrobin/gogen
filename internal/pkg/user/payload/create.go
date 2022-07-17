package payload

import (
	"github.com/paulusrobin/gogen-golib/mandatory"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type CreateUser struct {
	Mandatory  mandatory.Request
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (cu CreateUser) ToUserModel() model.User {
	return model.User{
		FirstName:  cu.FirstName,
		MiddleName: cu.MiddleName,
		LastName:   cu.LastName,
	}
}

func DecodeCreateRequest() {

}