package main

import (
	"log"
	"net/http"
	"todo-app/controllers"
	"todo-app/db"
	"todo-app/middleware"
	"todo-app/routes"
)

func main() {
	// Connect to MongoDB
	db.Connect()

	// Initialize router
	r := routes.SetupRouter()

	// Define public routes
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Define protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Setup CORS and RateLimiter middleware
	handler := middleware.RateLimiter(middleware.CORS(r))

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
