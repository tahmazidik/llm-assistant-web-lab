package models

import (
	"time"
)

type DialogID string

// Dialog отдельный диалог пользователя
type Dialog struct {
	ID       DialogID  `json:"dialog_id"`
	UserID   UserID    `json:"user_id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
