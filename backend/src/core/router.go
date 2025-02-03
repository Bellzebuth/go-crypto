package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// r.POST("/address/add", AddAddress)
	// r.DELETE("/address/:id", DeleteAddress)
	r.GET("/address/list", List)

	r.GET("/assets/search", SearchAsset)

	r.GET("/transactions/listsum", ListSum)
	r.GET("/transactions/list", List)

	// r.POST("/portfolio/dca", SimulateDCAHandler)

	// r.GET("/portfolio/charts", GetPortfolioChart)

	return r
}
