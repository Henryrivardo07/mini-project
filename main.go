package main

import (
	"log"
	"mini-project/routes"
	
)

func main() {
	// Setup Router
	r := routes.SetupRouter() // Menggunakan gin untuk setup router

	// Jalankan server pada port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
