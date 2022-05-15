package model

import (
	"time"
)

type ProductType struct {
	id          int       `json:"id" gorm:"column:id"`
	name        string    `json:"name" gorm:"column:name"`
	code        string    `json:"code" gorm:"column:code"`
	description string    `json:"description" gorm:"column:description"`
	price       float32      `json:"price" gorm:"column:price"`
	qty         float32      `json:"qty" gorm:"column:qty"`
	created_at  time.Time `json:"created_at" gorm:"column:created_at"`
	updated_at  time.Time `json:"update_at" gorm:"column:update_at"`
}
