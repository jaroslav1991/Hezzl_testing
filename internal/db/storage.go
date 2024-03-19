package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	MustBegin() *sqlx.Tx
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Storage struct {
	Db DB
}

func NewStorage(db DB) *Storage {
	return &Storage{Db: db}
}
