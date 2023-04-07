package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Username  string         `json:"username"`
	Email     string         `json:"email" gorm:"unique"`
	Phone     string         `json:"phone"`
	Password  string         `json:"password" validate:"required,min=6"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
