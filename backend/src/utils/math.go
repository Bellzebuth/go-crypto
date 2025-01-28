package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func FloatToInt(input float64, precision int) int {
	scale := math.Pow(10, float64(precision))
	return int(math.Round(input * scale))
}

func FormatPrecision(val float64, precision, digits int) (float64, error) {
	valWithPrecision := float64(val) / math.Pow(10, float64(precision))
	stringValWithDigits := fmt.Sprintf("%.*f", digits, valWithPrecision)
	return strconv.ParseFloat(stringValWithDigits, 64)
}

func CalculateGain(initialInvestment, purchasePrice, actualPrice int) (float64, float64, int, error) {
	if purchasePrice == 0 {
		return 0, 0, 0, errors.New("division by zero")
	}

	investment := float64(initialInvestment * 1_000_000)

	quantity := investment / float64(purchasePrice)
	currentValue := quantity * float64(actualPrice)

	gain := currentValue - investment
	percentageGain := (gain * 100) / investment

	formattedGain, err := FormatPrecision(gain, 6, 2)
	if err != nil {
		return 0, 0, 0, err
	}

	totalValue := float64(initialInvestment) + formattedGain

	return totalValue, formattedGain, int(percentageGain), nil
}
