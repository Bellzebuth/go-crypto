package controllers

import (
	"net/http"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func ListBlockchain(c *gin.Context) {
	var blockchains []models.Blockchain
	err := db.DB.Model(&blockchains).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blockchains)
}
