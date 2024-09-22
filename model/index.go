package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	UserID    uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoModel struct {
	TodoID    int64
	Todo      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
