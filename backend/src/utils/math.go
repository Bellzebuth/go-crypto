package utils

import (
	"math"
)

func FloatToInt(input float64, precision int) int {
	scale := math.Pow(10, float64(precision))
	return int(math.Round(input * scale))
}

func FormatPrecision(val, precision, virgule int) float64 {
	scale := math.Pow(10, float64(virgule))
	valWithDecimal := float64(val) / scale

	roundedValue := math.Round(valWithDecimal*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))

	return roundedValue
}

func CalculateGain(initialInvestment, buyPrice, newPrice int) (float64, int) {
	if buyPrice == 0 {
		return 0, 0 // avoid division by zero
	}

	quantity := initialInvestment * 1_000_000 / buyPrice
	currentValue := quantity * newPrice

	gain := currentValue - initialInvestment*1_000_000
	percentageGain := (gain * 100) / (initialInvestment * 1_000_000)

	return FormatPrecision(gain, 6, 2), percentageGain
}
