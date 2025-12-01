package memory

import (
	"context"
	"sync"

	models2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
)

// MessageRepository - in-memory хранилище сообщений
type MessageRepository struct {
	mu         sync.RWMutex
	byDialogID map[models2.DialogID][]*models2.Message
}

var _ dialogssvc.MessageRepository = (*MessageRepository)(nil)

// NewMessageRepository создает пустое хранилище сообщений
func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		byDialogID: make(map[models2.DialogID][]*models2.Message),
	}
}

// Create сохраняет новое сообщение
func (repo *MessageRepository) Create(ctx context.Context, message *models2.Message) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.byDialogID[message.DialogID] = append(repo.byDialogID[message.DialogID], message)

	return nil
}

func (repo *MessageRepository) ListByDialog(ctx context.Context, dialogID models2.DialogID) ([]*models2.Message, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	messages, ok := repo.byDialogID[dialogID]
	if !ok {
		return []*models2.Message{}, nil
	}

	return messages, nil
}
