package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bellzebuth/go-crypto/src/api"
	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func ListSum(c *gin.Context) {
	addressId, err := strconv.Atoi(c.Query("addressId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid addressId"})
		return
	}

	var totals []models.Transaction
	err = db.DB.Model(&totals).
		Relation("Address").
		Relation("Price").
		Relation("Price.Asset").
		ColumnExpr("SUM(value) AS value").
		ColumnExpr("SUM(purchased_price * value) / NULLIF(SUM(value), 1) AS purchased_price").
		Where("address_id = ?", addressId).
		Group("price__asset.name", "address_id", "address.id", "price.id", "price__asset.id").
		Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// for i, asset := range totals {
	// 	computedAsset, err := asset.ComputeGain()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	totals[i] = computedAsset
	// }

	c.JSON(http.StatusOK, totals)
}

func List(c *gin.Context) {
	var transactions []models.Transaction
	err := db.DB.Model(&transactions).
		Relation("Address").
		Relation("Price").
		Relation("Price.Asset").
		Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var computedAssets []models.Transaction
	for _, transaction := range transactions {
		computedAsset, err := transaction.ComputeGain()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		computedAssets = append(computedAssets, computedAsset)
	}

	c.JSON(http.StatusOK, computedAssets)
}

func LoadTransactions(address models.Address) error {
	transactions, err := api.GetTransactions(address)
	if err != nil {
		return err
	}

	if len(transactions) > 0 {
		_, err := db.DB.Model(&transactions).
			OnConflict("(id) DO NOTHING").
			Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateTransactions(userId int) error {
	addresses, err := listaddresses(userId)
	if err != nil {
		return err
	}

	// first execution
	fmt.Println("downloading transactions…")
	for _, address := range addresses {
		err = LoadTransactions(address)
		if err != nil {
			return err
		}
	}
	fmt.Println("transactions downloaded")

	return nil
}

// func GetTotals(c *gin.Context) {
// 	rows, err := db.DB.Query(`SELECT cp.key_name, a.amount, a.purchased_price, cp.price
// 		FROM assets AS a
// 		LEFT JOIN cache_prices AS cp
// 		ON a.key_name=cp.key_name`)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	totalValue := 0.0
// 	totalInvested := 0.0
// 	for rows.Next() {
// 		var asset Asset
// 		if err := rows.Scan(&asset.KeyName, &asset.Amount, &asset.PurchasedPrice, &asset.ActualPrice); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		totalInvested += asset.Amount

// 		value, _, _, err := utils.CalculateGain(asset.Amount, asset.PurchasedPrice, asset.ActualPrice)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		totalValue += value
// 	}

// 	c.JSON(http.StatusOK, gin.H{"totalInvested": totalInvested, "totalValue": fmt.Sprintf("%.2f", totalValue)})
// }
