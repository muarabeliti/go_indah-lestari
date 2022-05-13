package model

import (
	"time"

	"gorm.io/gorm"
)

type detail_resep struct {
	gorm.Model
	id_resep    int       	`json:"id_resep" gorm:"column:id_resep"`
	nm_resep  	float32  	`json:"nm_resep	" gorm:"column:nm_resep	"`
	cara_buat float32  	`json:"cara_buat" gorm:"cara_buat"`
	
}




