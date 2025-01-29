package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "crypto.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db

	log.Println("No tables found. Initializing database from schema.sql...")
	err = executeSQLFromFile(db, "schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}
	log.Println("Database initialized successfully.")

	return db, nil
}

func executeSQLFromFile(db *sql.DB, filename string) error {
	script, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filename, err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		return fmt.Errorf("failed to execute script: %v", err)
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
