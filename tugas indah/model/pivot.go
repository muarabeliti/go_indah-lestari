package model

import "time"

type Pivot struct {
	ID           int       `json:"id" gorm:"column:id"`
	IDProduct    int       `json:"id_product" gorm:"column:id_product"`
	IDLogProduct int       `json:"id_log_product" gorm:"column:id_log_product"`
	Qty          float32   `json:"qty" gorm:"column:qty"`
	Price        float32   `json:"price" gorm:"column:price"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}
