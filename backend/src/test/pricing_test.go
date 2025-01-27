package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCryptoPrice(t *testing.T) {
	// Mock API externe
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"bitcoin":{"usd":45000}}`))
	}))
	defer server.Close()

	originalURL := cryptoAPIURL
	cryptoAPIURL = server.URL
	defer func() { cryptoAPIURL = originalURL }()

	price, err := GetCryptoPrice("bitcoin")
	assert.NoError(t, err)
	assert.Equal(t, 45000.0, price)
}

func TestGetCryptoPriceInvalidAsset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	originalURL := cryptoAPIURL
	cryptoAPIURL = server.URL
	defer func() { cryptoAPIURL = originalURL }()

	_, err := GetCryptoPrice("invalid_asset")
	assert.Error(t, err)
}
