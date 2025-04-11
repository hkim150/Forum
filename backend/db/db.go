package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDb() (*sql.DB, error) {
	dbms := os.Getenv("DBMS")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open(dbms, dbUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
