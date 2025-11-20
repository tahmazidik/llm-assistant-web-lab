package models

import (
	"time"
)

type UserID string

// User описывает пользователя системы
type User struct {
	ID           UserID    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
}
