package user

import (
	"github.com/paulusrobin/gogen/internal/repository/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository function to initialize repository.
func NewRepository(db *gorm.DB) postgres.UserRepository {
	return repository{db: db}
}
