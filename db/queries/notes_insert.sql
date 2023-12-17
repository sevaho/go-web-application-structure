-- name: InsertNote :one

INSERT INTO notes (created_at, id, title, text)
    VALUES ($1, $2, $3, $4)
RETURNING
    serial_id;

