package dto

import "github.com/paulusrobin/gogen/internal/repository/model"

type (
	CreateRequest struct {
		User *model.User
	}
	CreateResponse struct {
		User *model.User
	}
)
