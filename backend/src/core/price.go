package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/utils"
)

var priceURL = "https://api.coingecko.com/api/v3/simple/price"

func buildURL() (string, error) {
	rows, err := db.DB.Query(`SELECT key_name FROM cryptos`)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return "", err
		}

		ids = append(ids, id)
	}

	return fmt.Sprintf("%s?ids=%s&vs_currencies=eur", priceURL, strings.Join(ids, ",")), nil
}

func UpdateCryptoPrices() error {
	url, err := buildURL()
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch price: %w", err)
	}

	defer resp.Body.Close()

	now := time.Now()
	if resp.StatusCode == 200 {
		var result map[string]map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return fmt.Errorf("failed to parse price response: %w", err)
		}

		for keyName, currencies := range result {
			for _, price := range currencies {
				// update cache
				query := `INSERT INTO cache_prices (key_name, price, last_update) VALUES (?, ?, ?)
    					ON CONFLICT(key_name) DO UPDATE SET
        					price = excluded.price,
        					last_update = excluded.last_update`

				_, err = db.DB.Exec(query, keyName, utils.ConvertToMicroUnits(price), now)
				if err != nil {
					return fmt.Errorf("failed to insert price for %s: %w", keyName, err)
				}
			}
		}
	} else {
		return fmt.Errorf("failed request with status code %d", resp.StatusCode)
	}

	return nil
}
