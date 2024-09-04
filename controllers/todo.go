package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo-app/db"
	"todo-app/models"
	"todo-app/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, http.StatusOK, "API Test successful")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	cursor, err := db.GetCollection().Find(context.Background(), bson.D{{}})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		todos = append(todos, todo)
	}
	utils.SendResponse(w, http.StatusOK, todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var todo models.Todo
	err = db.GetCollection().FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Todo not found")
		return
	}

	utils.SendResponse(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		log.Printf("CreateTodo - Decode Error: %v", err)
		return
	}

	// Basic validation
	if todo.Title == "" {
		utils.SendError(w, http.StatusBadRequest, "Title is required")
		log.Printf("CreateTodo - Validation Error: Title is required")
		return
	}

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	_, err = db.GetCollection().InsertOne(context.Background(), todo)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create todo")
		log.Printf("CreateTodo - Insert Error: %v", err)
		return
	}

	// Log successful creation
	log.Printf("CreateTodo - Success: %+v", todo)

	utils.SendResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "Todo created successfully",
		"todo":    todo,
	})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = db.GetCollection().UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": todo})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	_, err = db.GetCollection().DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(w, http.StatusOK, "Todo deleted successfully")
}
