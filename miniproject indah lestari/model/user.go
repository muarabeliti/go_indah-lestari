package model

import (
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	ID       int     `json:"id" gorm:"column:id"`
	username float32 `json:"username	" gorm:"column:username	"`
	password float32 `json:"password" gorm:"password"`
}
