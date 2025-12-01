package models

import (
	"time"
)

type MessageID string

// SenderType описывает, кто отправил сообщение
type SenderType string

const (
	SenderUser      SenderType = "user"
	SenderAssistant SenderType = "assistant"
)

// Message - сообщение в рамках диалога
type Message struct {
	ID       MessageID  `json:"id"`
	DialogID DialogID   `json:"dialog_id"`
	Sender   SenderType `json:"sender"`
	Content  string     `json:"content"`
	CreateAt time.Time  `json:"create_at"`
}
