package models

import (
	"time"
)

type TaskID string

// TaskStatus описывает статус задачи
type TaskStatus string

const (
	TaskStatusTODO       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

// Task - задача пользователя
type Task struct {
	ID          TaskID     `json:"id"`
	UserID      UserID     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
}
