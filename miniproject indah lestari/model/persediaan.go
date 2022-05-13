package model

import (
	"time"

	"gorm.io/gorm"
)

type persediaan struct {
	gorm.Model
	id    int       	`json:"id" gorm:"column:id"`
	id_user int 	`json:"id_user	" gorm:"column:id_user	"`
	bahan float32  	`json:"bahan" gorm:"bahan"`
	
}




