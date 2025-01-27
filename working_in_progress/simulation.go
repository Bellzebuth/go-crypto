package api

import (
	"net/http"
	"time"

	"github.com/Bellzebuth/go-crypto/src/services"
	"github.com/gin-gonic/gin"
)

func SimulateDCAHandler(c *gin.Context) {
	type Request struct {
		Asset      string  `json:"asset"`
		Investment float64 `json:"investment"`
		StartDate  string  `json:"start_date"`
		EndDate    string  `json:"end_date"`
	}

	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date"})
		return
	}

	result, err := services.SimulateDCA(req.Asset, req.Investment, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"final_value_usd": result})
}
