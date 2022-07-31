package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
}

func (User) TableName() string {
	return "users"
}
