package postgres

import (
	"context"
	"database/sql"

	models2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
)

type DialogRepository struct {
	db *sql.DB
}

var _ dialogssvc.DialogRepository = (*DialogRepository)(nil)

func NewDialogRepository(db *sql.DB) *DialogRepository {
	return &DialogRepository{db: db}
}

func (r *DialogRepository) Create(ctx context.Context, dialog *models2.Dialog) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO dialogs (id, user_id, title, create_at, update_at)
		VALUES ($1, $2, $3, $4, $5)
	`,
		string(dialog.ID),
		string(dialog.UserID),
		dialog.Title,
		dialog.CreateAt,
		dialog.UpdateAt,
	)
	return err
}

func (r *DialogRepository) GetByID(ctx context.Context, dialogID models2.DialogID) (*models2.Dialog, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, title, create_at, update_at
		FROM dialogs
		WHERE id = $1
	`, string(dialogID))

	var d models2.Dialog
	var id, uid string

	if err := row.Scan(&id, &uid, &d.Title, &d.CreateAt, &d.UpdateAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	d.ID = models2.DialogID(id)
	d.UserID = models2.UserID(uid)
	return &d, nil
}

func (r *DialogRepository) ListByUser(ctx context.Context, userID models2.UserID) ([]*models2.Dialog, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, title, create_at, update_at
		FROM dialogs
		WHERE user_id = $1
		ORDER BY update_at DESC
	`, string(userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models2.Dialog
	for rows.Next() {
		var d models2.Dialog
		var id, uid string
		if err := rows.Scan(&id, &uid, &d.Title, &d.CreateAt, &d.UpdateAt); err != nil {
			return nil, err
		}
		d.ID = models2.DialogID(id)
		d.UserID = models2.UserID(uid)
		out = append(out, &d)
	}
	return out, rows.Err()
}

func (r *DialogRepository) Update(ctx context.Context, dialog *models2.Dialog) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE dialogs
		SET title = $2, update_at = $3
		WHERE id = $1
	`, string(dialog.ID), dialog.Title, dialog.UpdateAt)
	return err
}

func (r *DialogRepository) Delete(ctx context.Context, dialogID models2.DialogID) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM dialogs
		WHERE id = $1
	`, string(dialogID))
	return err
}
