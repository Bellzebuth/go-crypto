package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bellzebuth/go-crypto/src/api"
	"github.com/stretchr/testify/assert"
)

func TestPortfolioWorkflowIntegration(t *testing.T) {
	router := api.SetupRouter()

	payload := []byte(`{"user_id":1,"name":"bitcoin","amount":0.1}`)
	req, _ := http.NewRequest("POST", "/portfolio/add", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest("GET", "/portfolio/value?user_id=1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"total_value_usd":4500`)

	req, _ = http.NewRequest("GET", "/portfolio/chart?user_id=1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Portfolio Performance")
}
