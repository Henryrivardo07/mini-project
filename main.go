package main

import (
	"fmt"
	"log"
	"mini-project/database"
	"mini-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	database.InitDB()

	// Setup Gin Router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	} else {
		fmt.Println("Server running on http://localhost:8080")
	}
}
