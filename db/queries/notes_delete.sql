-- name: DeleteNote :one

DELETE FROM
    notes n
WHERE
    n.id = $1
RETURNING
    n.serial_id;
