package models

import (
	"time"
)

type Article struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `form:"title" json:"title"`
	Description1 string    `form:"description1" json:"description1"`
	Description2 string    `form:"description2" json:"description2"`
	Description3 string    `form:"description3" json:"description3"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
