package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/Bellzebuth/go-crypto/src/utils"
)

var coinGeckoURL = "https://api.coingecko.com/api/v3/"

func buildCoinGeckoURL() (string, error) {
	var assets []string
	err := db.DB.Model(&models.Asset{}).
		Column("id").
		Select(&assets)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%ssimple/price?ids=%s&vs_currencies=eur", coinGeckoURL, strings.Join(assets, ",")), nil
}

func UpdateCryptoPrices() error {
	url, err := buildCoinGeckoURL()
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

type CoinGeckoPricesResponse struct {
	Prices [][]float64 `json:"prices"`
}

func GetEthereumPricesAt(timestamps []int64) (map[int64]float64, error) {
	if len(timestamps) == 0 {
		return nil, fmt.Errorf("timestamps is empty")
	}

	start := timestamps[0]
	end := timestamps[0]
	for _, ts := range timestamps {
		if ts < start {
			start = ts
		}
		if ts > end {
			end = ts
		}
	}

	// extends time slot to avoid errors
	start -= 60
	end += 60

	url := fmt.Sprintf("%scoins/ethereum/market_chart/range?vs_currency=eur&from=%d&to=%d", coinGeckoURL, start, end)

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CoinGeckoPricesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Prices) == 0 {
		return nil, fmt.Errorf("no price founded")
	}

	// sort prices by timestamps
	sort.Slice(result.Prices, func(i, j int) bool {
		return result.Prices[i][0] < result.Prices[j][0]
	})

	// match timestamps with prices
	pricesAtTimestamps := make(map[int64]float64)
	for _, ts := range timestamps {
		closestPrice := findClosestPrice(ts, result.Prices)
		pricesAtTimestamps[ts] = closestPrice
	}

	return pricesAtTimestamps, nil
}

func findClosestPrice(targetTimestamp int64, prices [][]float64) float64 {
	var closestPrice float64
	minDiff := math.MaxFloat64

	for _, entry := range prices {
		priceTimestamp := int64(entry[0] / 1000)
		price := entry[1]

		diff := math.Abs(float64(priceTimestamp - targetTimestamp))
		if diff < minDiff {
			minDiff = diff
			closestPrice = price
		}
	}

	return closestPrice
}
