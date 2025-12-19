package postgres

import (
	"context"
	"database/sql"

	models2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
)

type MessageRepository struct {
	db *sql.DB
}

var _ dialogssvc.MessageRepository = (*MessageRepository)(nil)

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(ctx context.Context, msg *models2.Message) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO messages (id, dialog_id, sender, content, create_at)
		VALUES ($1, $2, $3, $4, $5)
	`,
		string(msg.ID),
		string(msg.DialogID),
		string(msg.Sender),
		msg.Content,
		msg.CreateAt,
	)
	return err
}

func (r *MessageRepository) ListByDialog(ctx context.Context, dialogID models2.DialogID) ([]*models2.Message, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, dialog_id, sender, content, create_at
		FROM messages
		WHERE dialog_id = $1
		ORDER BY create_at ASC
	`, string(dialogID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models2.Message
	for rows.Next() {
		var m models2.Message
		var id, did, sender string

		if err := rows.Scan(&id, &did, &sender, &m.Content, &m.CreateAt); err != nil {
			return nil, err
		}

		m.ID = models2.MessageID(id)
		m.DialogID = models2.DialogID(did)
		m.Sender = models2.SenderType(sender)

		out = append(out, &m)
	}
	return out, rows.Err()
}

func (r *MessageRepository) DeleteByDialog(ctx context.Context, dialogID models2.DialogID) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM messages
		WHERE dialog_id = $1
	`, string(dialogID))
	return err
}
