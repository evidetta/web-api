package models

import (
	"database/sql"
	"time"
)

type DBModel interface {
	Update(db *sql.DB) error
	Delete(db *sql.DB) error
}

type Model struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
