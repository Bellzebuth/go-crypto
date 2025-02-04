package db

import (
	"log"

	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/go-pg/pg/v10/orm"
)

func MigrateDB() error {
	models := []interface{}{
		&models.User{},
		&models.Address{},
		&models.Asset{},
		&models.Price{},
		&models.Transaction{},
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	_, err := DB.Model(&AssetsList).
		OnConflict("DO NOTHING").
		Insert()
	if err != nil {
		return err
	}

	log.Println("✅ Migration réussie")
	return nil
}
