package user

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/model"
)

// Create function to create user object to database.
func (r repository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
