package controllers

import (
	"net/http"
	"mini-project/database"
	"mini-project/models"
	"mini-project/utils"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	rows, err := database.DB.Query("SELECT * FROM categories")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Error parsing data")
			return
		}
		categories = append(categories, category)
	}

	utils.SuccessResponse(c, http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	_, err := database.DB.Exec("INSERT INTO categories (name, created_at, created_by) VALUES ($1, NOW(), 'admin')", category.Name)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add category")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Category added successfully")
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	err := database.DB.QueryRow("SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category deleted successfully")
}
