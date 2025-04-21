package handlers

import (
	"encoding/json"
	"go-pvz/internal/auth"
	"go-pvz/internal/models"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user := models.User{
		Email:    userInput.Username,
		Password: userInput.Password,
		Role:     userInput.Role,
	}

	userID, err := auth.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"userID":  userID,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := auth.LoginUser(userInput.Username, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"token": token,
	}
	json.NewEncoder(w).Encode(response)
}
