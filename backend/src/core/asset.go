package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/utils"
	"github.com/gin-gonic/gin"
)

type Asset struct {
	Id             int       `json:"id"`
	KeyName        string    `json:"keyName"`
	Crypto         Crypto    `json:"crypto"`
	Amount         float64   `json:"amount"`
	PurchasedPrice float64   `json:"purchasedPrice"`
	CreatedAt      time.Time `json:"createdAt"`
	Gain           float64   `json:"gain"`
	PercentageGain float64   `json:"percentageGain"`
	ActualPrice    int       `json:"actualPrice"`
	ActualValue    float64   `json:"actualValue"`
}

func (a Asset) ComputeGain() (Asset, error) {
	value, gain, percentageGain, err := utils.CalculateGain(a.Amount, a.PurchasedPrice, a.ActualPrice)
	if err != nil {
		return a, err
	}

	a.Gain = gain
	a.PercentageGain = percentageGain
	a.ActualValue = value

	return a, nil
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

	query := `INSERT INTO assets (key_name, amount, purchased_price) VALUES (?, ?, ?)`
	_, err = db.DB.Exec(query, asset.KeyName, utils.ConvertToMicroUnits(asset.Amount), price.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "asset added"})
}

func ListSum(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT a.key_name, c.name, SUM(a.amount), AVG(purchased_price), cp.price
	FROM assets AS a 
	LEFT JOIN cryptos AS c 
	ON a.key_name = c.key_name
	LEFT JOIN cache_prices as cp
	ON a.key_name= cp.key_name
	GROUP BY a.key_name`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.KeyName, &asset.Crypto.Name, &asset.Amount, &asset.PurchasedPrice, &asset.ActualPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		asset, err := asset.ComputeGain()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		assets = append(assets, asset)
	}

	c.JSON(http.StatusOK, assets)
}

func List(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT a.id, c.key_name, c.name, a.amount, a.purchased_price, cp.price
	FROM assets AS a 
	LEFT JOIN cryptos AS c 
	ON a.key_name = c.key_name
	LEFT JOIN cache_prices as cp
	ON a.key_name= cp.key_name
	WHERE a.key_name = ?
	ORDER BY a.created_at DESC`, c.Query("keyName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.Id, &asset.KeyName, &asset.Crypto.Name, &asset.Amount, &asset.PurchasedPrice, &asset.ActualPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		asset, err := asset.ComputeGain()
		if err != nil {
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
	rows, err := db.DB.Query(`SELECT cp.key_name, a.amount, a.purchased_price, cp.price 
		FROM assets AS a 
		LEFT JOIN cache_prices AS cp 
		ON a.key_name=cp.key_name`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	totalValue := 0.0
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.KeyName, &asset.Amount, &asset.PurchasedPrice, &asset.ActualPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		value, _, _, err := utils.CalculateGain(asset.Amount, asset.PurchasedPrice, asset.ActualPrice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		totalValue += value
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%.2f", totalValue))
}
