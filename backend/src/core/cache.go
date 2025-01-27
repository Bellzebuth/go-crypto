package core

import (
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
)

type CachePrice struct {
	Id         int
	KeyName    int
	Crypto     *Crypto
	Price      int
	LastUpdate time.Time
}

func GetCachePrice(keyName string) (CachePrice, error) {
	var cachePrice CachePrice
	query := `SELECT * FROM cache_prices WHERE key_name = ?`

	err := db.DB.QueryRow(query, keyName).Scan(&cachePrice)
	if err != nil {
		return cachePrice, nil
	}

	return cachePrice, nil
}

// // PriceCacheEntry représente une entrée dans le cache avec sa valeur et son expiration
// type PriceCacheEntry struct {
// 	Price      float64
// 	Expiration time.Time
// }

// // PriceCache gère un cache concurrentiel pour les prix des cryptomonnaies
// type PriceCache struct {
// 	mu       sync.RWMutex
// 	cache    map[string]PriceCacheEntry
// 	lifetime time.Duration
// }

// // NewPriceCache crée une nouvelle instance de PriceCache
// func NewPriceCache(lifetime time.Duration) *PriceCache {
// 	return &PriceCache{
// 		cache:    make(map[string]PriceCacheEntry),
// 		lifetime: lifetime,
// 	}
// }

// // GetPrice récupère le prix depuis le cache ou effectue une requête externe si expiré
// func (pc *PriceCache) GetPrice(asset string, fetchPriceFunc func(string) (float64, error)) (float64, error) {
// 	pc.mu.RLock()
// 	entry, exists := pc.cache[asset]
// 	pc.mu.RUnlock()

// 	if exists && time.Now().Before(entry.Expiration) {
// 		// Le prix est encore valide dans le cache
// 		return entry.Price, nil
// 	}

// 	// Sinon, on doit récupérer le prix depuis l'API
// 	price, err := fetchPriceFunc(asset)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Stocker le nouveau prix dans le cache
// 	pc.mu.Lock()
// 	pc.cache[asset] = PriceCacheEntry{
// 		Price:      price,
// 		Expiration: time.Now().Add(pc.lifetime),
// 	}
// 	pc.mu.Unlock()

// 	return price, nil
// }
