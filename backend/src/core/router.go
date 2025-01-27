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

	r.POST("/portfolio/add", Add)
	r.GET("/cryptos/list", Search)
	r.GET("/portfolio/list", List)
	r.DELETE("/portfolio/:id", Delete)
	r.GET("/portfolio/total", GetTotal)

	// r.POST("/portfolio/dca", SimulateDCAHandler)

	// r.GET("/portfolio/charts", GetPortfolioChart)

	return r
}
