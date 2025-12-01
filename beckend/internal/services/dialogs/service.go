package dialogs

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	models2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
)

// Понятные ошибки доменного уровня для диалогов
var (
	ErrDialogNotFound      = errors.New("dialog not found")
	ErrEmptyDialogTitle    = errors.New("dialog title is empty")
	ErrEmptyMessageContent = errors.New("message content is empty")
)

// Service описывает бизнес-логику для работы с диалогами
type Service interface {
	// CreateDialog создает новый диалог для пользователя
	CreateDialog(ctx context.Context, userID models2.UserID, title string) (*models2.Dialog, error)

	//ListDialogs возвращает список диалогов пользователя
	ListDialogs(ctx context.Context, userID models2.UserID) ([]*models2.Dialog, error)

	// AddMessage добавляет сообщение в диалог
	AddMessage(ctx context.Context, dialogID models2.DialogID, sender models2.SenderType, content string) (*models2.Message, error)

	// ListMessages возвращает все сообщения диалога
	ListMessages(ctx context.Context, dialogID models2.DialogID) ([]*models2.Message, error)
}

// DialogRepository описывает операции с диалогами в хранилище
type DialogRepository interface {
	Create(ctx context.Context, dialogID *models2.Dialog) error
	GetByID(ctx context.Context, dialogID models2.DialogID) (*models2.Dialog, error)
	ListByUser(ctx context.Context, userID models2.UserID) ([]*models2.Dialog, error)
	Update(ctx context.Context, dialog *models2.Dialog) error
}

// MessageRepository описывает операции с сообщениями в хранилище
type MessageRepository interface {
	Create(ctx context.Context, msg *models2.Message) error
	ListByDialog(ctx context.Context, dialogID models2.DialogID) ([]*models2.Message, error)
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

func (svc *service) CreateDialog(ctx context.Context, userID models2.UserID, title string) (*models2.Dialog, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrEmptyDialogTitle
	}

	now := svc.now()

	dialog := &models2.Dialog{
		ID:       models2.DialogID(uuid.NewString()),
		UserID:   userID,
		Title:    title,
		CreateAt: now,
		UpdateAt: now,
	}

	if err := svc.dialogRepo.Create(ctx, dialog); err != nil {
		return nil, err
	}

	return dialog, nil
}

func (svc *service) ListDialogs(ctx context.Context, userID models2.UserID) ([]*models2.Dialog, error) {
	return svc.dialogRepo.ListByUser(ctx, userID)
}

func (svc *service) AddMessage(ctx context.Context, dialogID models2.DialogID, sender models2.SenderType, content string) (*models2.Message, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, ErrEmptyMessageContent
	}

	// Проверяем что диалог существует
	dialog, err := svc.dialogRepo.GetByID(ctx, dialogID)
	if err != nil {
		return nil, err
	}

	if dialog == nil {
		return nil, ErrDialogNotFound
	}

	now := svc.now()

	msg := &models2.Message{
		ID:       models2.MessageID(uuid.NewString()),
		DialogID: dialogID,
		Sender:   sender,
		Content:  content,
		CreateAt: now,
	}

	// Сохраняем сообщение
	if err := svc.messageRepo.Create(ctx, msg); err != nil {
		return nil, err
	}

	// Обновляем время обновления диалога
	dialog.UpdateAt = now
	if err := svc.dialogRepo.Update(ctx, dialog); err != nil {
		return nil, err
	}

	return msg, nil
}

// Возвращает все сообщения диалога
func (svc *service) ListMessages(ctx context.Context, dialogID models2.DialogID) ([]*models2.Message, error) {
	//TODO: проверить что диалог существует
	return svc.messageRepo.ListByDialog(ctx, dialogID)
}
