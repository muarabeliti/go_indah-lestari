package model

import "gorm.io/gorm"

type UserLogin struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" gorm:"column:password"`
	Token    string `json:"token" form:"token"`
}
