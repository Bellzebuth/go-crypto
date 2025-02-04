package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/Bellzebuth/go-crypto/src/utils"
)

var priceURL = "https://api.coingecko.com/api/v3/simple/price"

func buildURL() (string, error) {
	var assets []string
	err := db.DB.Model(&models.Asset{}).
		Column("id").
		Select(&assets)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s?ids=%s&vs_currencies=eur", priceURL, strings.Join(assets, ",")), nil
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

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed request with status code %d", resp.StatusCode)
	}

	var result map[string]map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to parse price response: %w", err)
	}

	var prices []models.Price
	for keyName, currencies := range result {
		for _, price := range currencies {
			prices = append(prices, models.Price{
				AssetId:    keyName,
				Price:      utils.ConvertToMicroUnits(price),
				LastUpdate: now,
			})
		}
	}

	if len(prices) > 0 {
		_, err = db.DB.Model(&prices).
			OnConflict("(asset_id) DO UPDATE").
			Set("price = excluded.price, last_update = excluded.last_update").
			Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
