package memory

import (
	"context"
	"sync"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	usersvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
)

// UserRepository - простое in-memory хранилище пользователей
// Хранит все в map, защищенной мьютексом
type UserRepository struct {
	mu      sync.RWMutex
	byID    map[models.UserID]*models.User
	byEmail map[string]*models.User
}

// Complete-time проверка, что мы реализуем интерфейс usersvc.Repository
// Можно ли значение типа *UserRepository положить в переменную типа usersvc.Repository?
// А по правилам Go можно положить, только если:
// тип *UserRepository реализует интерфейс usersvc.Repository
// (то есть имеет все его методы с нужными подписями).
var _ usersvc.Repository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		byID:    make(map[models.UserID]*models.User),
		byEmail: make(map[string]*models.User),
	}
}

// Create сохраняет нового пользователя
func (ur *UserRepository) Create(ctx context.Context, user *models.User) error {
	// TODO: можно проверять ctx.Done()
	ur.mu.Lock()
	defer ur.mu.Unlock()

	ur.byID[user.ID] = user
	ur.byEmail[user.Email] = user

	return nil
}

// GetByEmail ищет пользователя по email
func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	user, ok := ur.byEmail[email]
	if !ok {
		return nil, nil
	}

	return user, nil
}

// GetByID ищет пользователя по ID
func (ur *UserRepository) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	user, ok := ur.byID[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}
