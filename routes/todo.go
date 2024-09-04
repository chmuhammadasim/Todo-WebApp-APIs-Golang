package routes

import (
	"todo-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/test", controllers.Test).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	return r
}
