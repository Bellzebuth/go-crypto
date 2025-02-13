package main

import (
	"log"
	"os"
	"time"

	"github.com/Bellzebuth/go-crypto/src/api"
	"github.com/Bellzebuth/go-crypto/src/controllers"
	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/joho/godotenv"
)

func updatecryptoPrices() {
	log.Println("Starting price update...")
	err := api.UpdateCryptoPrices()
	if err != nil {
		log.Printf("Error updating crypto prices: %v", err)
	} else {
		log.Println("Crypto prices updated successfully.")
	}
}

func resetDBAndStart() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = db.ResetDB()
	if err != nil {
		return err
	}

	err = startServer()
	if err != nil {
		return err
	}

	return nil
}

func startServer() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// db connexion
	err = db.ConnectDB()
	if err != nil {
		return err
	}

	// close connexion at the end
	defer db.CloseDB()

	// do migration
	err = db.MigrateDB()
	if err != nil {
		return err
	}

	r := controllers.SetupRouter()

	go func() {
		// first execution
		updatecryptoPrices()

		// execute every 10 minutes
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			updatecryptoPrices()
		}
	}()

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <command>\nAvailable commands: reset, server")
	}

	command := os.Args[1]

	switch command {
	case "reset":
		err := resetDBAndStart()
		if err != nil {
			panic(err)
		}
	case "server":
		err := startServer()
		if err != nil {
			panic(err)
		}
	default:
		log.Fatalf("Unknown command: %s\nAvailable commands: reset, server", command)
	}
}
