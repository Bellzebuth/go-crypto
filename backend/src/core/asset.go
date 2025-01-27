package core

import (
	"net/http"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/utils"
	"github.com/gin-gonic/gin"
)

type Asset struct {
	Id          int       `json:"id"`
	KeyName     string    `json:"keyName"`
	Crypto      Crypto    `json:"crypto"`
	Amount      int       `json:"amount"`
	BuyingPrice int       `json:"buyingPrice"`
	CreatedAt   time.Time `json:"createdAt"`
}

func Add(c *gin.Context) {
	var asset Asset
	err := c.ShouldBindJSON(&asset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price, err := GetCachePrice(asset.KeyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO assets (key_name, amount, buying_price) VALUES (?, ?)`
	_, err = db.DB.Exec(query, asset.KeyName, asset.Amount, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "asset added"})
}

func List(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT a.id, c.key_name, c.name, a.amount, a.created_at FROM assets AS a LEFT JOIN cryptos AS c WHERE a.key_name = c.key_name`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.Id, &asset.Crypto.KeyName, &asset.Crypto.Name, &asset.Amount, &asset.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		assets = append(assets, asset)
	}

	c.JSON(http.StatusOK, assets)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(`DELETE FROM assets WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Crypto deleted"})
}

func GetTotal(c *gin.Context) {
	rows, err := db.DB.Query("SELECT cp.key_name, a.amount, a.buying_price, cp.price FROM assets AS a LEFT JOIN cache_prices AS cp ON a.key_name=cp.key_name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	totalValue := 0.0
	for rows.Next() {
		asset := struct {
			keyName     string
			amount      int
			buyingPrice int
			price       int
		}{}

		if err := rows.Scan(&asset.keyName, &asset.amount, &asset.buyingPrice, &asset.price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		gain, _ := utils.CalculateGain(asset.amount, asset.buyingPrice, asset.price)
		totalValue += gain
	}

	c.JSON(http.StatusOK, totalValue)
}
