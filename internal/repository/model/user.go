package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey"`
	FirstName  string
	MiddleName string
	LastName   string
}
