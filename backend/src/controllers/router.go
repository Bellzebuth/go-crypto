package controllers

import (
	"github.com/Bellzebuth/go-crypto/src/middleware"
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

	r.POST("/register", Register)
	r.POST("/login", Login)
	r.GET("/logout", Logout)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/protected", ProtectedRoute)

	// auth.POST("/address/add", AddAddress)
	// auth.DELETE("/address/:id", DeleteAddress)
	auth.POST("/address/add", AddAddress)
	auth.GET("/address/list", ListAddress)

	auth.GET("/blockchain/list", ListBlockchain)

	auth.GET("/assets/search", SearchAsset)

	auth.GET("/transactions/listsum", ListSum)
	auth.GET("/transactions/list", List)

	// auth.POST("/portfolio/dca", SimulateDCAHandler)

	// auth.GET("/portfolio/charts", GetPortfolioChart)

	return r
}
