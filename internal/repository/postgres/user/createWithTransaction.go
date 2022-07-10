package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"gorm.io/gorm"
)

// CreateWithTransaction function to create user object to database using transaction.
func (r repository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error {
	return tx.WithContext(ctx).Create(user).Error
}
