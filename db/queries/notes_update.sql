-- name: UpdateNoteBySerialID :one
UPDATE
    notes n
SET
    title = CASE WHEN @title_do_update::boolean
        THEN @title::TEXT ELSE title END,
    text = CASE WHEN @text_do_update::boolean
        THEN @text::TEXT ELSE text END
WHERE
    n.id = @note_id
RETURNING
    n.serial_id;
