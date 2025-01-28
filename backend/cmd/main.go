package main

import (
	"log"
	"os"
	"time"

	"github.com/Bellzebuth/go-crypto/src/core"
	"github.com/Bellzebuth/go-crypto/src/db"
)

func resetDB() {
	log.Println("Resetting the database...")

	err := db.ResetDB()
	if err != nil {
		log.Fatalf("Failed to reset the database: %v", err)
	}

	log.Println("Database reset successfully.")
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()
	defer db.CloseDB()

	r := core.SetupRouter()

	go func() {
		log.Println("Starting price update...")
		err := core.UpdateCryptoPrices() // first execution
		if err != nil {
			log.Printf("Error updating crypto prices: %v", err)
		} else {
			log.Println("Crypto prices updated successfully.")
		}

		// execute every 10 minutes
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("Update cache prices ...")
			err := core.UpdateCryptoPrices()
			if err != nil {
				log.Printf("Error updating crypto prices: %v", err)
			} else {
				log.Println("Crypto prices updated successfully.")
			}
		}
	}()

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <command>\nAvailable commands: reset_db, server")
	}

	command := os.Args[1]

	switch command {
	case "reset_db":
		resetDB()
	case "server":
		startServer()
	default:
		log.Fatalf("Unknown command: %s\nAvailable commands: reset_db, server", command)
	}
}
