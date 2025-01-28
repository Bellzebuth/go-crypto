package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "crypto.db"
const schemaFile = "schema.sql"

// ResetDB supprime et recrée la base de données à partir de schema.sql
func ResetDB() error {
	// Supprimer la base de données existante
	if _, err := os.Stat(dbFile); err == nil {
		err = os.Remove(dbFile)
		if err != nil {
			return fmt.Errorf("failed to remove existing database: %v", err)
		}
	}

	// Créer une nouvelle base vide
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	defer db.Close()

	// Charger le schéma SQL
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Exécuter le schéma
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil
}
