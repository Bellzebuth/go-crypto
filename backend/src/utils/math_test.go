package utils_test

import (
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
		val       int
		precision int
		virgule   int
		expected  float64
	}{
		{12345, 3, 2, 123.45},         // Décalage virgule à 2 et appliquer 3 chiffres après
		{987654, 6, 3, 0.987},         // Décalage virgule à 3 et appliquer 6 chiffres après
		{1, 1, 2, 0.10},               // Décalage virgule à 2 et appliquer 1 chiffre après
		{5000, 3, 2, 5.000},           // Décalage virgule à 2 et appliquer 3 chiffres après
		{999999, 6, 2, 0.999999},      // Décalage virgule à 2 et appliquer 6 chiffres après
		{123456789, 8, 2, 1.23456789}, // Décalage virgule à 2 et appliquer 8 chiffres après
	}

	for _, tt := range tests {
		t.Run("Test FormatPrecision", func(t *testing.T) {
			result := utils.FormatPrecision(tt.val, tt.precision, tt.virgule)
			assert.Equal(t, tt.expected, result, "Expected value does not match")
		})
	}
}

func TestCalculateGain(t *testing.T) {
	tests := []struct {
		initialInvestment      int
		buyPrice               int
		newPrice               int
		expectedGain           float64
		expectedPercentageGain int
	}{
		// Cas basiques
		{1000, 500, 700, 200.0, 20},   // Gain de 200 sur un investissement initial de 1000
		{2000, 1000, 1500, 500.0, 25}, // Gain de 500 sur un investissement initial de 2000

		// Cas avec un gain nul
		{1500, 1000, 1000, 0.0, 0}, // Aucun gain
		{1000, 1000, 1000, 0.0, 0}, // Aucun gain, prix inchangé

		// Cas où le prix d'achat est nul (éviter division par zéro)
		{1000, 0, 1500, 0.0, 0}, // Division par zéro, donc gain = 0

		// Cas où la valeur de l'investissement est inférieure au prix d'achat
		{1000, 2000, 1000, -1000.0, -50}, // Perte de 1000 sur un investissement initial de 1000

		// Cas où la valeur de l'investissement est supérieure au prix d'achat
		{2000, 1000, 500, -500.0, -25}, // Perte de 500 sur un investissement initial de 2000
	}

	for _, tt := range tests {
		t.Run("Test CalculateGain", func(t *testing.T) {
			resultGain, resultPercentageGain := utils.CalculateGain(tt.initialInvestment, tt.buyPrice, tt.newPrice)
			assert.Equal(t, tt.expectedGain, resultGain, "Expected gain does not match")
			assert.Equal(t, tt.expectedPercentageGain, resultPercentageGain, "Expected percentage gain does not match")
		})
	}
}
