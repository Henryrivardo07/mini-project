package controllers

import (
	"github.com/gin-gonic/gin"
	"mini-project/database"
	"mini-project/models"
	"net/http"
	"strconv"
)

// GetCategoryByID mendapatkan kategori berdasarkan ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	err = database.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}
