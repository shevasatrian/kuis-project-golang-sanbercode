package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
