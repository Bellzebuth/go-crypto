package controllers

import (
	"net/http"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func AddAddress(c *gin.Context) {
	user, err := GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	address := models.Address{
		UserId: int(user.Id),
	}

	err = c.ShouldBindJSON(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err = db.DB.Model(&address).Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't insert address"})
		return
	}

	c.JSON(http.StatusOK, address)
}

func ListAddress(c *gin.Context) {
	user, err := GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var address []models.Address
	err = db.DB.Model(&address).
		Where("userId = ?", user.Id).
		Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, address)
}

func DeleteAddress(c *gin.Context) {
	user, err := GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var address string

	err = c.ShouldBindJSON(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err = db.DB.Model(&models.Address{}).
		Where("user_id = ?", user.Id).
		Where("address = ?", address).
		Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't insert address"})
		return
	}

	c.JSON(http.StatusOK, address)
}
