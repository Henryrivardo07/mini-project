package controllers

import (
	"net/http"
	"time"
	"mini-project/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SecretKey = []byte("your_secret_key")

func Login(c *gin.Context) {
	var loginData map[string]string
	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	username, password := loginData["username"], loginData["password"]
	if username != "admin" || password != "password" { // Sederhanakan untuk demo
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error generating token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"token": tokenString})
}
