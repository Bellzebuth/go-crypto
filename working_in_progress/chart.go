package api

import (
	"net/http"

	"github.com/Bellzebuth/go-crypto/src/services"
	"github.com/gin-gonic/gin"
)

func GetPortfolioChart(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	dates := []string{"2023-01-01", "2023-02-01", "2023-03-01", "2023-04-01"}
	values := []float64{1000, 1200, 1500, 1800}

	outputPath := "./portfolio_chart.html"
	if err := services.GeneratePortfolioChart(dates, values, outputPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate chart"})
		return
	}

	c.File(outputPath)
}
