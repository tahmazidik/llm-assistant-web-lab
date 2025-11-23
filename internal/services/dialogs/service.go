package dialogs

import (
	"context"
	"time"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
)

// Service описывает бизнес-логику для работы с диалогами
type Service interface {
	// CreateDialog создает новый диалог для пользователя
	CreateDialog(ctx context.Context, userID models.UserID, title string) (*models.Dialog, error)

	//ListDialogs возвращает список диалогов пользователя
	ListDialogs(ctx context.Context, userID models.UserID) ([]*models.Dialog, error)

	// AddMessage добавляет сообщение в диалог
	AddMessage(ctx context.Context, dialogID models.DialogID, sender models.SenderType, content string) (*models.MessageID, error)

	// ListMessages возвращает все сообщения диалога
	ListMessages(ctx context.Context, dialogID models.DialogID) ([]*models.Message, error)
}

// DialogRepository описывает операции с даалогами в хранилище
type DialogRepository interface {
	Create(ctx context.Context, dialogID models.Dialog) error
	GetByID(ctx context.Context, dialogID models.DialogID) (*models.Dialog, error)
	ListByUser(ctx context.Context, userID models.UserID) ([]*models.Dialog, error)
	Update(ctx context.Context, dialog *models.Dialog) error
}

// MessageRepository описывает операции с сообщениями в хранилище
type MessageRepository interface {
	Create(ctx context.Context, dialogID models.Dialog) error
	ListByUser(ctx context.Context, userID models.UserID) ([]*models.Dialog, error)
}

// service - конкретная реализация Service
type service struct {
	dialogRepo  DialogRepository
	messageRepo MessageRepository
	now         func() time.Time
}

func NewService(dialogRepo DialogRepository, messageRepo MessageRepository) Service {
	return &service{
		dialogRepo:  dialogRepo,
		messageRepo: messageRepo,
		now:         time.Now,
	}
}

func (s *service) CreateDialog(ctx context.Context, userID models.UserID, title string) (*models.Dialog, error) {
	return nil, nil
}

func (s *service) ListDialogs(ctx context.Context, userID models.UserID) ([]*models.Dialog, error) {
	return nil, nil
}

func (s *service) AddMessage(ctx context.Context, dialogID models.DialogID, sender models.SenderType, content string) (*models.MessageID, error) {
	return nil, nil
}

func (s *service) ListMessages(ctx context.Context, dialogID models.DialogID) ([]*models.Message, error) {
	return nil, nil
}
