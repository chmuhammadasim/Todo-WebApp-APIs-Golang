package main

import (
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/routes"

	"github.com/rs/cors"
)

func main() {
	// Connect to MongoDB
	db.Connect()

	// Initialize router
	r := routes.SetupRouter()

	// Setup CORS
	handler := cors.Default().Handler(r)

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
