package postgres

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error
}
