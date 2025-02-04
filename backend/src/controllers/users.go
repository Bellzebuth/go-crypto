package controllers

import (
	"net/http"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// hash password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	_, err := db.DB.Model(&user).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// get user
	var user models.User
	err := db.DB.Model(&user).Where("username = ?", creds.Username).Select()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// check password
	if !user.CheckPassword(creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// set cookie
	c.SetCookie("session_token", creds.Username, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
}

func Logout(c *gin.Context) {
	// delete cookie
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("session_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// check user
		var user models.User
		err = db.DB.Model(&user).Where("username = ?", username).Select()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		// add user to context
		c.Set("user", user)
		c.Next()
	}
}

func ProtectedRoute(c *gin.Context) {
	user, _ := c.Get("user") // get user from middleware
	c.JSON(http.StatusOK, gin.H{"message": "Welcome", "user": user})
}
