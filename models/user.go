package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username     string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	FirstName    string    `json:"first_name" gorm:"type:varchar(100)"`
	LastName     string    `json:"last_name" gorm:"type:varchar(100);not null"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	PasswordHash string    `json:"password_hash" gorm:"type:varchar(255);not null"`
	Role         string    `json:"role" gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
