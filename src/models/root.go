package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type TaskCreateCmd struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type TaskQueryModel struct {
	Title string `json:"title"`
	Text  string `json:"text"`

	SerialID  int       `json:"-"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
