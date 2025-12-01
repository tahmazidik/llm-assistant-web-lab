package models

import (
	"time"
)

type NoteID struct{}

// Note - заметка пользователя
type Note struct {
	ID       NoteID    `json:"id"`
	TaskID   TaskID    `json:"task_id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
