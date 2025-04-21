package handlers

import (
	"encoding/json"
	"go-pvz/internal/auth"
	"go-pvz/internal/models"
	"net/http"
)

// RegisterHandler — обработчик для регистрации пользователя
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	// Декодируем тело запроса
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Создаем структуру User для передачи в функцию RegisterUser
	user := models.User{
		Email:    userInput.Username,
		Password: userInput.Password,
		Role:     userInput.Role,
	}

	// Вызов функции регистрации пользователя
	userID, err := auth.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ответ с сообщением об успешной регистрации
	response := map[string]interface{}{
		"message": "User registered successfully",
		"userID":  userID,
	}
	json.NewEncoder(w).Encode(response)
}

// LoginHandler — обработчик для логина пользователя
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Декодируем тело запроса
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Вызов функции логина и получение токена
	token, err := auth.LoginUser(userInput.Username, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Ответ с токеном
	response := map[string]interface{}{
		"token": token,
	}
	json.NewEncoder(w).Encode(response)
}
