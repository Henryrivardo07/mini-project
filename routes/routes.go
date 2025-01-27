package routes

import (
	"mini-project/controllers"
	"mini-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/api/categories")
	categoryRoutes.GET("/", controllers.GetCategories)
	categoryRoutes.POST("/", middlewares.JWTAuth(), controllers.AddCategory)
	categoryRoutes.GET("/:id", controllers.GetCategoryByID)
	categoryRoutes.DELETE("/:id", middlewares.JWTAuth(), controllers.DeleteCategory)

	bookRoutes := router.Group("/api/books")
	bookRoutes.GET("/", controllers.GetBooks)
	bookRoutes.POST("/", middlewares.JWTAuth(), controllers.AddBook)
	bookRoutes.DELETE("/:id", middlewares.JWTAuth(), controllers.DeleteBook)
}
