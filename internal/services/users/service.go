package users

import (
	"context"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
)

// Service описывает бизнес-логику для работы с пользователями
type Service interface {
	// Register регистрирует нового пользователя
	Register(ctx context.Context, email, password, name string) (*models.User, error)

	// Authenticate проверяет email и пароль
	// Возвращает пользователя, если пара логин/пароль верна
	Authenticate(ctx context.Context, email, password string) (*models.User, error)

	// GetByID возвращает пользователя по его ID
	GetByID(ctx context.Context, id models.UserID) (*models.User, error)
}

type Repository interface {
	// Create создает нового пользователя
	Create(ctx context.Context, user *models.User) error

	// GetByEmail находит пользователя по email
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// GetByID находит пользователя по ID
	GetByID(ctx context.Context, id models.UserID) (*models.User, error)
}
