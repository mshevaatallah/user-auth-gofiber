package response

import (
	"gorm.io/gorm"
	"time"
)

type UserResponse struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Username  string         `json:"username"`
	Email     string         `json:"email" gorm:"unique"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
