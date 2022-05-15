package model

import (
	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" gorm:"column:password"`
}
