package core

import (
	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
)

func GetPrice(keyName string) (models.Price, error) {
	var price models.Price
	err := db.DB.Model(&price).
		Relation("Asset").
		Where("asset.key_name = ?", keyName).
		Select()

	return price, err
}
