package core

import (
	"net/http"
	"strings"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func SearchAsset(c *gin.Context) {
	query := strings.ToLower(c.Query("query"))

	assets := []models.Asset{}
	err := db.DB.Model(&assets).
		Where("name ILIKE ?", "%"+query+"%").
		Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assets)
}
