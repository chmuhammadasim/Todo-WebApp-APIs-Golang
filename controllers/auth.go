// controllers/auth.go
package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-app/db"
	"todo-app/models"
	"todo-app/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key") // Use a more secure key in production

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user models.User
	err = db.GetCollection().FindOne(context.Background(), map[string]string{"username": loginData.Username}).Decode(&user)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID.Hex(),
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	utils.SendResponse(w, http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}
