package model

import (
	"time"
)

type Post struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Content string ` gorm:"type:text" json:"content" `
	Category string `json:"category"`
	Status string `'json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMPS" json:"created_at"`
	UpdatedAT time.Time `gorm:"default:CURRENT_TIMESTAMPS" json:"updated_at"`
}
