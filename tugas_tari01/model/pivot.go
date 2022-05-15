package model

import "time"

type Pivot struct {
	idlogproduct   int       `json:"id_log_product" gorm:"column:id_log_product"`
	qty            float32       `json:"qty" gorm:"column:qty"`
	price          int       `json:"price" gorm:"column:price"`
	id_log_product int       `json:"id_product" gorm:"column:id_product"`
	created_at     time.Time `json:"created_at" gorm:"column:created_at"`
	updated_at     time.Time `json:"update_at" gorm:"column:update_at"`
}
