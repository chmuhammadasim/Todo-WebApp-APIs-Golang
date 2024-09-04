package main

import (
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/middleware"
	"todo-app/routes"
)

func main() {
	// Connect to MongoDB
	db.Connect()

	// Initialize router
	r := routes.SetupRouter()

	// Setup CORS and RateLimiter middleware
	handler := middleware.CORS(middleware.RateLimiter(r))

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
