package core

import (
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
)

type CachePrice struct {
	Id         int
	KeyName    string
	Price      int
	LastUpdate time.Time
}

func GetCachePrice(keyName string) (CachePrice, error) {
	var cachePrice CachePrice
	query := `SELECT id, key_name, price, last_update FROM cache_prices WHERE key_name = ?`

	err := db.DB.QueryRow(query, keyName).Scan(&cachePrice.Id, &cachePrice.KeyName, &cachePrice.Price, &cachePrice.LastUpdate)

	return cachePrice, err
}
