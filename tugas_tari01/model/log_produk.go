package model

import "time"

type logproduct struct {
	id 			int			`json:"id" gorm:"column:id"`
	status		int			`json:"status" gorm:"column:status"`
	total_qty	int			`json:"total_qty" gorm:"column:total_qty"`
	created_at	time.Time	`json:"created_at" gorm:"column:created_at"`
	updated_at	time.Time	`json:"updated_at" gorm:"column:updated_at"`

}