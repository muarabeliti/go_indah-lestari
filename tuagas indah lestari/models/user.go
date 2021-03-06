package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"-" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
