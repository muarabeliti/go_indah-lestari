package model

import (
	"time"

	"gorm.io/gorm"
)

type resep struct {
	gorm.Model
	id    int       	`json:"id" gorm:"column:id"`
	id_bahan int 	`json:"id_bahan	" gorm:"column:id_bahan	"`
	nama float32  	`json:"nama" gorm:"nama"`
	deskripsi float32  	`json:"deskripsi" gorm:"deskripsi"`
	
}




