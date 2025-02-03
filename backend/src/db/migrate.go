package db

import (
	"log"

	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/go-pg/pg/v10/orm"
)

func MigrateDB() {
	models := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatalf("Erreur de migration: %v", err)
		}
	}

	log.Println("✅ Migration réussie")
}
