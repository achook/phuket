package database

import (
	"database/sql"
	"fmt"
	"os"

	"phuket/logging"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (d *Database) connect(dbConfig string) error {
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return err
	}

	d.db = db
	return nil
}

func (d *Database) close() error {
	return d.db.Close()
}

func (d *Database) Query(query string) (*sql.Rows, error) {
	return d.db.Query(query)
}

func (d *Database) QueryRow(query string) *sql.Row {
	return d.db.QueryRow(query)
}

func (d *Database) Exec(query string) (sql.Result, error) {
	return d.db.Exec(query)
}

func NewDatabase() (*Database, error) {
	d := new(Database)

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		logging.LogFatalMessage("POSTGRES_USER is not set")
	}

	dbName := os.Getenv("POSTGRES_DB_NAME")
	if dbName == "" {
		logging.LogFatalMessage("POSTGRES_DB_NAME is not set")
	}

	conf := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbName)
	if err := d.connect(conf); err != nil {
		return d, err
	}

	return d, nil
}
