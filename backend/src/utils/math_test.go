package utils_test

import (
	"errors"
	"testing"

	"github.com/Bellzebuth/go-crypto/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		precision int
		expected  int
	}{
		{
			name:      "Simple decimal",
			input:     0.292431,
			precision: 6,
			expected:  292431,
		},
		{
			name:      "Rounding up",
			input:     0.2924319,
			precision: 6,
			expected:  292432,
		},
		{
			name:      "Rounding down",
			input:     0.2924311,
			precision: 6,
			expected:  292431,
		},
		{
			name:      "Different precision",
			input:     0.292431,
			precision: 4,
			expected:  2924,
		},
		{
			name:      "Zero input",
			input:     0.0,
			precision: 6,
			expected:  0,
		},
		{
			name:      "Negative input",
			input:     -0.292431,
			precision: 6,
			expected:  -292431,
		},
		{
			name:      "Large precision",
			input:     0.292431,
			precision: 8,
			expected:  29243100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.FloatToInt(tt.input, tt.precision)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatPrecision(t *testing.T) {
	tests := []struct {
		val       float64
		precision int
		virgule   int
		expected  float64
	}{
		{12345, 3, 2, 12.35},
		{987654, 6, 3, 0.988},
		{1, 1, 2, 0.10},
		{5000, 3, 2, 5.00},
		{999999, 6, 2, 1.00},
		{123456789, 8, 2, 1.23},
	}

	for _, tt := range tests {
		t.Run("Test FormatPrecision", func(t *testing.T) {
			result, _ := utils.FormatPrecision(tt.val, tt.precision, tt.virgule)
			assert.Equal(t, tt.expected, result, "Expected value does not match")
		})
	}
}

func TestCalculateGain(t *testing.T) {
	tests := []struct {
		initialInvestment      int
		buyPrice               int
		newPrice               int
		expectedTotal          float64
		expectedGain           float64
		expectedPercentageGain int
		expectedErr            error
	}{
		// Cas basiques
		{20, 23500000, 23500000, 20.0, 0.0, 0, nil},
		{1000, 500, 700, 1400.0, 400.0, 40, nil},
		{2000, 1000, 1500, 3000.0, 1000.0, 50, nil},

		// Cas avec un gain nul
		{1500, 1000, 1000, 1500.0, 0.0, 0, nil}, // Aucun gain
		{1000, 1000, 1000, 1000.0, 0.0, 0, nil}, // Aucun gain, prix inchangé

		// Cas où le prix d'achat est nul (éviter division par zéro)
		{1000, 0, 1500, 0.0, 0.0, 0, errors.New("division by zero")}, // Division par zéro, donc gain = 0

		// Cas où la valeur de l'investissement est inférieure au prix d'achat
		{1000, 2000, 1000, 500.0, -500.0, -50, nil}, // Perte de 1000 sur un investissement initial de 1000

		// Cas où la valeur de l'investissement est supérieure au prix d'achat
		{2000, 1000, 500, 1000.0, -1000.0, -50, nil},
	}

	for _, tt := range tests {
		t.Run("Test CalculateGain", func(t *testing.T) {
			resultTotal, resultGain, resultPercentageGain, err := utils.CalculateGain(tt.initialInvestment, tt.buyPrice, tt.newPrice)
			assert.Equal(t, tt.expectedTotal, resultTotal, "Expected total does not match")
			assert.Equal(t, tt.expectedGain, resultGain, "Expected gain does not match")
			assert.Equal(t, tt.expectedPercentageGain, resultPercentageGain, "Expected percentage gain does not match")
			assert.Equal(t, tt.expectedErr, err, "Expected err does not match")
		})
	}
}
