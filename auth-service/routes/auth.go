package routes

import (
	"net/http"
	"todo-app/auth-service/models"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]models.User)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", register)
	router.POST("/login", login)
}

func register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users[user.Username] = user
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := users[user.Username]; !exists || users[user.Username].Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
