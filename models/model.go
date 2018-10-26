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
	Tag       string     `json:"tag,omitempty"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
