package db

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func ConnectDB() error {
	dbName := "crypto"
	user := "bellzebuth"
	password := "password" // ENV VARIABLE
	addr := "localhost:5432"

	// connection to postgres
	tempDB := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: "postgres",
	})

	exists, err := databaseExists(tempDB, dbName)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Println("📦 Database creation…")
		err := createDatabase(tempDB, dbName)
		if err != nil {
			return err
		}
		fmt.Println("✅ Database created")
	}

	tempDB.Close()

	// connection to the main database
	DB = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbName,
	})

	// check connection
	if err := DB.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("✅ Connected")
	return nil
}

func databaseExists(db *pg.DB, dbName string) (bool, error) {
	var exists bool
	_, err := db.QueryOne(pg.Scan(&exists), "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)", dbName)
	return exists, err
}

func createDatabase(db *pg.DB, dbName string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	return err
}

func CloseDB() {
	DB.Close()
}
