package services

import (
	"fmt"
	"time"
)

func SimulateDCA(asset string, investment float64, startDate, endDate time.Time) (float64, error) {
	historicalPrices := map[string]float64{
		"2023-01-01": 50000,
		"2023-02-01": 45000,
		"2023-03-01": 40000,
	}

	var totalInvested, totalCoins float64
	for date, price := range historicalPrices {
		investmentDate, err := time.Parse("2006-01-02", date)
		if err != nil || investmentDate.Before(startDate) || investmentDate.After(endDate) {
			continue
		}

		totalInvested += investment
		totalCoins += investment / price
	}

	finalPrice, err := GetCryptoPrice(asset)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch final price: %w", err)
	}

	return totalCoins * finalPrice, nil
}
