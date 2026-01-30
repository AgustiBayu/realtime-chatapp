package repository

import (
	"context"
	"database/sql"
	"realtime-chatapp/internal/domain"
)

type MessageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.MessageRepository {
	return &MessageRepositoryImpl{
		db: db,
	}
}

func (r *MessageRepositoryImpl) Save(ctx context.Context, message domain.Message) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	SQL := "INSERT INTO messages (sender_id, receiver_id, content, created_at) VALUES($1,$2,$3,$4)"
	if _, err := tx.ExecContext(ctx, SQL, message.SenderID, message.ReceiverID, message.Content, message.CreatedAt); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *MessageRepositoryImpl) GetChatHistory(ctx context.Context, senderID, receiverID int) ([]domain.MessageResponse, error) {
	SQL := `SELECT m.id, m.sender_id, m.receiver_id, m.content, m.created_at,
		       u1.name as sender_name, u2.name as receiver_name
		FROM messages m
		JOIN users u1 ON m.sender_id = u1.id
		JOIN users u2 ON m.receiver_id = u2.id
		WHERE (m.sender_id = $1 AND m.receiver_id = $2)
		   OR (m.sender_id = $2 AND m.receiver_id = $1)
		ORDER BY m.created_at ASC`

	rows, err := r.db.QueryContext(ctx, SQL, senderID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var history []domain.MessageResponse
	for rows.Next() {
		var msg domain.MessageResponse
		err := rows.Scan(
			&msg.ID,
			&msg.Sender.ID,
			&msg.Receiver.ID,
			&msg.Content,
			&msg.CreatedAt,
			&msg.Sender.Name,
			&msg.Receiver.Name,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, msg)
	}
	return history, nil
}
