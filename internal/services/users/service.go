package users

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
)

// Понятные ошибки доменного уровня
var (
	ErrEmailAlreadyInUse  = errors.New("email already in use")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
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

type service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
		now:  time.Now,
	}
}

func (s *service) Register(ctx context.Context, email, password, name string) (*models.User, error) {
	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	//Проверяем что такого email еще нет
	existing, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrEmailAlreadyInUse
	}

	//Хеэшируем пароль (bсrypt - стандартный вариант для Go)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	now := s.now()

	//Собираем модель пользователя
	user := &models.User{
		ID:           models.UserID(uuid.NewString()),
		Email:        email,
		PasswordHash: string(hashedPassword),
		Name:         name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	//Сохряняем в хранилище через репозиторий
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Authenticate проверяет email и пароль
func (s *service) Authenticate(ctx context.Context, email, password string) (*models.User, error) {
	email = strings.TrimSpace(email)

	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	// Сравниваем хэш пароль с переданным паролем
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *service) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
