package memory

import (
	"context"
	"sync"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
)

// MessageRepository - in-memory хранилище сообщений
type MessageRepository struct {
	mu         sync.RWMutex
	byDialogID map[models.DialogID][]*models.Message
}

var _ dialogssvc.MessageRepository = (*MessageRepository)(nil)

// NewMessageRepository создает пустое хранилище сообщений
func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		byDialogID: make(map[models.DialogID][]*models.Message),
	}
}

// Create сохраняет новое сообщение
func (repo *MessageRepository) Create(ctx context.Context, message *models.Message) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.byDialogID[message.DialogID] = append(repo.byDialogID[message.DialogID], message)

	return nil
}

func (repo *MessageRepository) ListByDialog(ctx context.Context, dialogID models.DialogID) ([]*models.Message, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	messages, ok := repo.byDialogID[dialogID]
	if !ok {
		return []*models.Message{}, nil
	}

	return messages, nil
}
