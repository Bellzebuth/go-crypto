package core

import (
	"testing"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/stretchr/testify/assert"
)

type MockPricingService struct{}

func (m *MockPricingService) GetCryptoPrice(asset string) (float64, error) {
	if asset == "bitcoin" {
		return 45000, nil
	}
	return 0, nil
}

func TestGetPortfolioValue(t *testing.T) {
	db.DB.Exec("INSERT INTO portfolios (user_id, name, amount) VALUES (1, 'bitcoin', 0.1)")

	value, err := GetPortfolioValue(1)
	assert.NoError(t, err)
	assert.Equal(t, 4500.0, value)
}
