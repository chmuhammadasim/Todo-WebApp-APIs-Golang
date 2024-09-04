// controllers/auth.go
package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-app/db"
	"todo-app/models"
	"todo-app/utils"

	"golang.org/x/crypto/bcrypt"
)

//var secretKey = []byte("your_secret_key") // Use a more secure key in production

func Signup(w http.ResponseWriter, r *http.Request) {
	var signupData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // e.g., "user" or "admin"
	}
	err := json.NewDecoder(r.Body).Decode(&signupData)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Check if user already exists
	var existingUser models.User
	err = db.GetCollection().FindOne(context.Background(), map[string]string{"username": signupData.Username}).Decode(&existingUser)
	if err == nil {
		utils.SendError(w, http.StatusConflict, "User already exists")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Create user
	user := models.User{
		Username: signupData.Username,
		Password: string(hashedPassword),
		Role:     signupData.Role,
	}
	_, err = db.GetCollection().InsertOne(context.Background(), user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendResponse(w, http.StatusCreated, "User created successfully")
}
