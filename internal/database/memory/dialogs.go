package memory

import (
	"context"
	"sync"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
)

// DialogRepository - in-memory хранилище диалогов
type DialogRepository struct {
	mu       sync.RWMutex
	byID     map[models.DialogID]*models.Dialog
	byUserID map[models.UserID][]*models.Dialog
}

var _ dialogssvc.DialogRepository = (*DialogRepository)(nil)

// NewDialogRepository создает пустое хранилище диалогов
func NewDialogRepository() *DialogRepository {
	return &DialogRepository{
		byID:     make(map[models.DialogID]*models.Dialog),
		byUserID: make(map[models.UserID][]*models.Dialog),
	}
}

// Create сохраняет новый диалог
func (dialogRepo *DialogRepository) Create(ctx context.Context, dialog *models.Dialog) error {
	dialogRepo.mu.Lock()
	defer dialogRepo.mu.Unlock()

	dialogRepo.byID[dialog.ID] = dialog
	dialogRepo.byUserID[dialog.UserID] = append(dialogRepo.byUserID[dialog.UserID], dialog)

	return nil
}

// GetByID ищет диалог по ID
func (dialogRepo *DialogRepository) GetByID(ctx context.Context, id models.DialogID) (*models.Dialog, error) {
	dialogRepo.mu.RLock()
	defer dialogRepo.mu.RUnlock()

	dialog, ok := dialogRepo.byID[id]
	if !ok {
		return nil, nil
	}

	return dialog, nil
}

// ListByUser возвращает все диалоги пользователя
func (dialogRepo *DialogRepository) ListByUser(ctx context.Context, userID models.UserID) ([]*models.Dialog, error) {
	dialogRepo.mu.RLock()
	defer dialogRepo.mu.RUnlock()

	dialogs, ok := dialogRepo.byUserID[userID]
	if !ok {
		return []*models.Dialog{}, nil
	}

	return dialogs, nil
}

// Update обновляет диалог
func (dialogRepo *DialogRepository) Update(ctx context.Context, dialog *models.Dialog) error {
	dialogRepo.mu.Lock()
	defer dialogRepo.mu.Unlock()

	// Обновляем по ID
	dialogRepo.byID[dialog.ID] = dialog

	//Предполагаем, что диалог уже есть в хранилище, если нет - просто перезапишем
	dialogs := dialogRepo.byUserID[dialog.UserID]
	updated := false
	for i, d := range dialogs {
		if d.ID == dialog.ID {
			dialogs[i] = dialog
			updated = true
			break
		}
	}

	if !updated {
		// если не нашли - добавим
		dialogs = append(dialogs, dialog)
	}

	dialogRepo.byUserID[dialog.UserID] = dialogs

	return nil
}
