package controllers

import (
	"net/http"
	"mini-project/database"
	"mini-project/models"
	"mini-project/utils"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	rows, err := database.DB.Query("SELECT * FROM books")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch books")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Error parsing data")
			return
		}
		books = append(books, book)
	}

	utils.SuccessResponse(c, http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Set thickness based on total page
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	_, err := database.DB.Exec("INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), 'admin')", book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add book")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Book added successfully")
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Book deleted successfully")
}
