package routes

import (
	"github.com/gin-gonic/gin"
	"mini-project/controllers"
	"mini-project/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Example of a public route
	r.GET("/public", controllers.PublicRoute)

	// Protected routes with JWT authentication
	protected := r.Group("/protected")
	protected.Use(middlewares.JWTAuth) // Apply JWTAuth middleware globally to this group
	{
		protected.GET("/category/:id", controllers.GetCategoryByID)
		protected.GET("/books", controllers.GetBooks)
	}

	return r
}
