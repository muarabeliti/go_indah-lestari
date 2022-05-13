package model

import (
	"gorm.io/gorm"
)

type bahan struct {
	gorm.Model
	ID          int     `json:"id" gorm:"column:id"`
	Protein     float32 `json:"Protein	" gorm:"column:Protein	"`
	karbohidrat float32 `json:"karbohidrat" gorm:"karbohidrat"`
	bumbu       float32 `json:"bumbu " gorm:"bumbu "`
}
