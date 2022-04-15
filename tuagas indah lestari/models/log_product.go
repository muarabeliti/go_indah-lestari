package models

import "time"

type LogProduct struct {
	ID        			int      	`json:"id" gorm:"column:id"`
	Status    			string    	`json:"status" gorm:"column:status"`
	TotalQty			float32   		`json:"total_qty" gorm:"column:total_qty"`
	TotalPrice			float32   	`json:"total_price" gorm:"column:total_price"`
	CreatedAt 			time.Time	`json:"created_at" gorm:"column:created_at"`
	UpdatedAt 			time.Time 	`json:"updated_at" gorm:"column:updated_at"`
}
