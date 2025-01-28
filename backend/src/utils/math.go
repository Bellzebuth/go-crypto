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

func CalculateGain(initialInvestment int, purchasePrice float64, actualPrice int) (float64, float64, int, error) {
	if purchasePrice == 0 {
		return 0, 0, 0, errors.New("division by zero")
	}

	investment := float64(initialInvestment * 1_000_000)

	quantity := investment / purchasePrice
	currentValue := quantity * float64(actualPrice)

	gain := currentValue - investment
	percentageGain := (gain * 100) / investment

	formattedGain, err := FormatPrecision(gain, 6, 2)
	if err != nil {
		return 0, 0, 0, err
	}

	totalValue := float64(investment) + formattedGain

	return totalValue, formattedGain, int(percentageGain), nil
}
