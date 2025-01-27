package main

import (
	"log"
	"os"
	"time"

	"github.com/Bellzebuth/go-crypto/src/core"
	"github.com/Bellzebuth/go-crypto/src/db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()
	defer db.CloseDB()

	r := core.SetupRouter()

	go func() {
		ticker := time.NewTicker(1 * time.Hour) // execute every hour
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Starting price update...")
				err := core.UpdateCryptoPrices()
				if err != nil {
					log.Printf("Error updating crypto prices: %v", err)
				} else {
					log.Println("Crypto prices updated successfully.")
				}
			}
		}
	}()

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
