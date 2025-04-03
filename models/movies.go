package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title     string    `gorm:"size:255;not null" json:"title" binding:"required"`
	Director  string    `gorm:"size:255;not null" json:"director" binding:"required"`
	Year      int       `json:"year" binding:"required"`
	Plot      string    `gorm:"type:text" json:"plot"`
	Genre     string    `gorm:"size:100" json:"genre"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
}
