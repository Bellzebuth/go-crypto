package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./crypto.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTables()
}

func createTables() {
	portfolioTable := `CREATE TABLE IF NOT EXISTS portfolios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		amount REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	simulationTable := `CREATE TABLE IF NOT EXISTS simulations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		strategy TEXT NOT NULL,
		result REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(portfolioTable); err != nil {
		log.Fatalf("Failed to create portfolios table: %v", err)
	}

	if _, err := DB.Exec(simulationTable); err != nil {
		log.Fatalf("Failed to create simulations table: %v", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
