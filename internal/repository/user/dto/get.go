package dto

import "github.com/paulusrobin/gogen/internal/repository/model"

type (
	GetByIDRequest struct {
		UserID uint
	}
	GetByIDResponse struct {
		User *model.User
	}
)
