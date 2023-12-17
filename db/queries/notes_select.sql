-- name: SelectNotes :many

SELECT
  *
FROM
    notes
ORDER BY
    created_at DESC;

