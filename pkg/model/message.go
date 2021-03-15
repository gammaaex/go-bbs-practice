package handler

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type MessageRecord struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type MessageRecords []*MessageRecord

type MessageTable struct{}

func (m *MessageTable) SelectAll(ctx context.Context, db *sql.DB) (MessageRecords, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM messages;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := make(MessageRecords, 0)
	for rows.Next() {
		record := &MessageRecord{}
		if err := rows.Scan(&record.ID, &record.Content); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (m *MessageTable) Insert(ctx context.Context, db *sql.DB, record *MessageRecord) error {
	record.ID = uuid.New().String()

	stmt, err := db.Prepare("INSERT INTO messages(id, content) VALUES(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, record.ID, record.Content)
	if err != nil {
		return err
	}

	return nil
}

func (m *MessageTable) Update(ctx context.Context, db *sql.DB, record *MessageRecord) error {
	stmt, err := db.Prepare("UPDATE messages SET content = ? WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, record.Content, record.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MessageTable) Delete(ctx context.Context, db *sql.DB, record *MessageRecord) error {
	stmt, err := db.Prepare("DELETE FROM messages WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, record.ID)
	if err != nil {
		return err
	}

	return nil
}
