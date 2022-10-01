package model

import (
	"github.com/paulusrobin/gogen-golib/encoding/json"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FirstName  string         `gorm:"first_name" json:"first_name"`
	MiddleName string         `gorm:"middle_name" json:"middle_name"`
	LastName   string         `gorm:"last_name" json:"last_name"`
	CreatedAt  time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}

func (u User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}
