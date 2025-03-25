package db

import (
	"database/sql"
)

// DBTX defines the interface for database operations
type DBTX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// New creates a new Queries instance
func New(db DBTX) *Queries {
	return &Queries{db: db}
}
