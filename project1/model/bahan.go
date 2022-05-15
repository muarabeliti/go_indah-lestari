package models

import "time"

type Bahan struct {
	Id          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IdUser      int    `json:"id_user" form:"id_user"`
	Ingredients string `json:"ingredients" form:"ingredients"`
}
