package models

import (
	"time"
)

// Photo represents the data model for a photo.
type Photo struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title"`
	Caption  string    `json:"caption"`
	PhotoURL string    `json:"photo_url"`
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
