package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestGetPortfolioValueHandler(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/portfolio/value", GetTotal)

// 	services.InitTestDB()
// 	services.DB.Exec("INSERT INTO portfolios (user_id, asset_name, amount) VALUES (1, 'bitcoin', 0.1)")

// 	req, _ := http.NewRequest("GET", "/portfolio/value?user_id=1", nil)
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	assert.Contains(t, resp.Body.String(), `"total_value_usd":4500`)
// }

func TestGetPortfolioValueHandlerInvalidUserID(t *testing.T) {
	router := gin.Default()
	router.GET("/portfolio/value", GetTotal)

	req, _ := http.NewRequest("GET", "/portfolio/value?user_id=invalid", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), `"error":"invalid user_id"`)
}

func TestGetPortfolioValueHandlerMissingUserID(t *testing.T) {
	router := gin.Default()
	router.GET("/portfolio/value", GetTotal)

	req, _ := http.NewRequest("GET", "/portfolio/value", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), `"error":"user_id is required"`)
}
