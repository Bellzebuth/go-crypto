package utils

import (
	"errors"
	"math"
)

func ConvertToMicroUnits(value float64) int64 {
	return int64(math.Floor(value * 1_000_000))
}

func CalculateGain(initialInvestment float64, purchasePrice float64, actualPrice int64) (float64, float64, float64, error) {
	if purchasePrice == 0 {
		return 0, 0, 0, errors.New("division by zero")
	}

	if purchasePrice == float64(actualPrice) {
		return initialInvestment, 0, 0, nil
	}

	quantity := initialInvestment / purchasePrice

	currentValue := quantity * float64(actualPrice)

	gain := currentValue - initialInvestment
	percentageGain := (gain * 100) / initialInvestment

	totalValue := initialInvestment + gain

	return totalValue, gain, percentageGain, nil
}
