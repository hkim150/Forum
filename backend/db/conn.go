package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbms     = "postgres"
	host     = "db"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "forum"
)

func Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open(dbms, psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
