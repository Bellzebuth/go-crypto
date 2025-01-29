package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func ConvertToMicroUnits(value float64) int64 {
	return int64(math.Floor(value * 1_000_000))
}

func FormatPrecision(val float64, precision, digits int) (float64, error) {
	valWithPrecision := float64(val) / math.Pow(10, float64(precision))
	stringValWithDigits := fmt.Sprintf("%.*f", digits, valWithPrecision)
	return strconv.ParseFloat(stringValWithDigits, 64)
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
