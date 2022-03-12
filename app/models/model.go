package models

import (
	"database/sql"
	"time"
)

type Model struct {
	ID uint `json:"id"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletedAt struct {
	DeletedAt sql.NullTime `json:"deleted_at"`
}
