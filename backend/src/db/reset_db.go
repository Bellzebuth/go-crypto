package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "crypto.db"
const schemaFile = "schema.sql"

func ResetDB() error {
	// delete existing database
	if _, err := os.Stat(dbFile); err == nil {
		err = os.Remove(dbFile)
		if err != nil {
			return fmt.Errorf("failed to remove existing database: %v", err)
		}
	}

	// create empty database
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	defer db.Close()

	// upload schema
	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// execute schema
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil
}
