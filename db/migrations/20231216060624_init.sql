-- migrate:up

CREATE TABLE notes (
    serial_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id uuid NOT NULL UNIQUE,
    created_at timestamp NOT NULL,
    title varchar(256) NOT NULL,
    text TEXT NOT NULL
);

-- migrate:down
