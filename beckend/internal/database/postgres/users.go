package postgres

import (
	"context"
	"database/sql"

	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	usersvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/users"
)

type UserRepository struct {
	db *sql.DB
}

var _ usersvc.Repository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	const q = `
        INSERT INTO users (id, email, name, password_hash, create_at, update_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.ExecContext(ctx, q,
		user.ID,
		user.Email,
		user.Name,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	const q = `
        SELECT id, email, name, password_hash, create_at, update_at
        FROM users
        WHERE email = $1
    `
	row := r.db.QueryRowContext(ctx, q, email)

	var u models.User
	if err := row.Scan(&u.ID, &u.Email, &u.Name, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	const q = `
        SELECT id, email, name, password_hash, create_at, update_at
        FROM users
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, q, id)

	var u models.User
	if err := row.Scan(&u.ID, &u.Email, &u.Name, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
