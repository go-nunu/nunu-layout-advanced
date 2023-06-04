package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func (u *User) TableName() string {
	return "users"
}
