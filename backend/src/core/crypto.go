package core

import (
	"net/http"
	"strings"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/gin-gonic/gin"
)

type Crypto struct {
	KeyName string `json:"keyName"`
	Name    string `json:"name"`
}

func GetCryptoByKeyName(keyName string) (Crypto, error) {
	var crypto Crypto

	query := `SELECT * FROM cryptos WHERE key_name = ?`
	err := db.DB.QueryRow(query, keyName).Scan(&crypto)
	if err != nil {
		return crypto, err
	}

	return crypto, nil
}

func Search(c *gin.Context) {
	query := strings.ToLower(c.Query("query"))

	cryptos := []Crypto{}
	rows, err := db.DB.Query(`SELECT * FROM cryptos where name LIKE ? COLLATE NOCASE LIMIT 5`, "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var crypto Crypto
		if err := rows.Scan(&crypto.KeyName, &crypto.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cryptos = append(cryptos, crypto)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cryptos)
}
