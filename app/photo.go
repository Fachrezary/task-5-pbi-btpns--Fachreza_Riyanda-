// app/photo.go

package app

import (
	"time"
)

// Photo struct represents the data model for a photo.
type Photo struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title" validate:"required"`
	Caption  string    `json:"caption"`
	PhotoURL string    `json:"photo_url" validate:"required,url"`
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
