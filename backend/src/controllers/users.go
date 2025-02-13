package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromRequest(c *gin.Context) (*models.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, errors.New("inexistant user")
	}

	// try cast
	user, ok := userInterface.(models.User)
	if !ok {
		return nil, errors.New("failed to cast user")
	}

	return &user, nil
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// hash password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	_, err := db.DB.Model(&user).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

var userCancelFuncs = make(map[int]context.CancelFunc)
var mu sync.Mutex // Pour éviter les conflits d'accès concurrentiel

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// get user
	var user models.User
	err := db.DB.Model(&user).Where("username = ?", creds.Username).Select()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// check password
	if !user.CheckPassword(creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())

	mu.Lock()
	userCancelFuncs[user.Id] = cancel
	mu.Unlock()

	go func() {
		err := UpdateTransactions(user.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update transactions"})
			return
		}

		// execute every hours
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := UpdateTransactions(user.Id)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update transactions"})
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// set cookie
	c.SetCookie("session_token", creds.Username, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
}

func Logout(c *gin.Context) {
	user, err := GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// delete user's go routine
	mu.Lock()
	if cancel, exists := userCancelFuncs[user.Id]; exists {
		cancel()
		delete(userCancelFuncs, user.Id)
	}
	mu.Unlock()

	// delete cookie
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func ProtectedRoute(c *gin.Context) {
	user, _ := c.Get("user") // get user from middleware
	c.JSON(http.StatusOK, gin.H{"message": "Welcome", "user": user})
}
