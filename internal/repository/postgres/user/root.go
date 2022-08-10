package user

import (
	repository2 "github.com/paulusrobin/gogen/internal/repository"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository function to initialize repository.
func NewRepository(db *gorm.DB) repository2.UserRepository {
	return repository{db: db}
}
