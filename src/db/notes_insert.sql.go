// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: notes_insert.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertNote = `-- name: InsertNote :one

INSERT INTO notes (created_at, id, title, text)
    VALUES ($1, $2, $3, $4)
RETURNING
    serial_id
`

type InsertNoteParams struct {
	CreatedAt pgtype.Timestamp
	ID        pgtype.UUID
	Title     string
	Text      string
}

func (q *Queries) InsertNote(ctx context.Context, arg InsertNoteParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertNote,
		arg.CreatedAt,
		arg.ID,
		arg.Title,
		arg.Text,
	)
	var serial_id int32
	err := row.Scan(&serial_id)
	return serial_id, err
}
