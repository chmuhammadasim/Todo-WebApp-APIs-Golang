// main.go
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

	// Define protected routes for users
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware("user")) // Role-based access for regular users
	api.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")

	// Define protected routes for admins
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware("admin")) // Role-based access for admins
	// Add admin routes here, e.g., managing users, etc.

	// Setup CORS and RateLimiter middleware
	handler := middleware.RateLimiter(middleware.CORS(r))

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
