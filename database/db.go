package database

import (
	"database/sql"
	"fmt"
	"os"

	"phuket/logging"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connect() error {

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		logging.LogFatalMessage("POSTGRES_USER is not set")
	}

	dbName := os.Getenv("POSTGRES_DB_NAME")
	if dbName == "" {
		logging.LogFatalMessage("POSTGRES_DB_NAME is not set")
	}

	conf := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbName)

	var err error
	db, err = sql.Open("postgres", conf)
	if err != nil {
		return err
	}

	return nil
}

func close() error {
	return db.Close()
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return db.Query(query, args...)
}

func QueryRow(query string, args ...any) *sql.Row {
	return db.QueryRow(query, args...)
}

func Exec(query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}
