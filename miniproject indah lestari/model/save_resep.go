package model

import (
	"time"

	"gorm.io/gorm"
)

type save_resep struct {
	gorm.Model
	id    int       	`json:"id" gorm:"column:id"`
	id_user int 	`json:"id_user	" gorm:"column:id_user"`
	id_resep int 	`json:"id_resep	" gorm:"column:id_resep"`
}




