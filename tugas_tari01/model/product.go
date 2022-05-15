package model

import "time"

type Product struct {
	ID        			int       `json:"id" gorm:"column:id"`
	Code      			string    `json:"code" gorm:"column:code"`
	ProductTypeId      	int       `json:"product_type_id" gorm:"column:product_type_id"`
	Name      			string    `json:"name" gorm:"column:name"`
	Description      	string    `json:"description" gorm:"column:description"`
	Price				float32   `json:"price" gorm:"column:price"`
	Qty					float32		  `json:"qty" gorm:"column:qty"`
	CreatedAt 			time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt 			time.Time `json:"updated_at" gorm:"column:updated_at"`
}
