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
		log.Println("Starting price update...")
		err := core.UpdateCryptoPrices() // initial execution
		if err != nil {
			log.Printf("Error updating crypto prices: %v", err)
		} else {
			log.Println("Crypto prices updated successfully.")
		}

		// Re execute every 10 minutes
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Update cache prices ...")
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
