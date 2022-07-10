package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, user model.User) error
		CreateWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error
	}
	repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}
