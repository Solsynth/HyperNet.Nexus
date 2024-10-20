package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Kdb *sql.DB

func Connect(str string) (*sql.DB, error) {
	db, err := sql.Open("postgres", str)
	if err == nil {
		Kdb = db
	}
	return db, err
}
