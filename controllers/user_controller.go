package controllers

import (
	"mini-project/database"
	"mini-project/models" // Pastikan ini ada
	"mini-project/utils"
	"github.com/gin-gonic/gin"
)

// GetUser retrieves user data
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	err := database.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		utils.ErrorResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, "Invalid input")
		return
	}

	_, err := database.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		utils.ErrorResponse(c, "Failed to create user")
		return
	}

	utils.SuccessResponse(c, "User created successfully")
}
