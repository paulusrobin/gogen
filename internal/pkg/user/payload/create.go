package payload

import (
	"github.com/paulusrobin/gogen/internal/repository/model"
)

type CreateUser struct {
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
