package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublicRoute adalah handler untuk route publik
func PublicRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the public route!",
	})
}
