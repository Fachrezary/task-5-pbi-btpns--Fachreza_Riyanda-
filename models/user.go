package models

import (
	"time"
)

// User represents the data model for a user.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}
