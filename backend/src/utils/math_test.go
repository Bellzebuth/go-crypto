package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToMicroUnits(t *testing.T) {
	tests := []struct {
		input    float64
		expected int64
	}{
		{1, 1000000},
		{2.34, 2340000},
		{1.23456789, 1234567},
		{0.000001, 1},
		{123.456789, 123456789},
		{0, 0},
		{123456.789, 123456789000},
	}

	for _, test := range tests {
		t.Run("Testing ConvertToMicroUnits", func(t *testing.T) {
			result := ConvertToMicroUnits(test.input)
			if result != test.expected {
				t.Errorf("Pour l'entrée %.8f, attendu %d, obtenu %d", test.input, test.expected, result)
			}
		})
	}
}

func TestCalculateGain(t *testing.T) {
	tests := []struct {
		initialInvestment      float64
		PurchasedPrice         float64
		newPrice               int64
		expectedTotal          float64
		expectedGain           float64
		expectedPercentageGain float64
		expectedErr            error
	}{
		// Cas basiques
		{200000000, 98331000000, 98331000000, 200000000, 0.0, 0, nil},
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
			resultTotal, resultGain, resultPercentageGain, err := CalculateGain(tt.initialInvestment, tt.PurchasedPrice, tt.newPrice)
			assert.Equal(t, tt.expectedTotal, resultTotal, "Expected total does not match")
			assert.Equal(t, tt.expectedGain, resultGain, "Expected gain does not match")
			assert.Equal(t, tt.expectedPercentageGain, resultPercentageGain, "Expected percentage gain does not match")
			assert.Equal(t, tt.expectedErr, err, "Expected err does not match")
		})
	}
}
