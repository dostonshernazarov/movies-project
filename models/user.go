package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
