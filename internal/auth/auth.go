package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/romangergovskiy/go-pvz/internal/database"
	"github.com/romangergovskiy/go-pvz/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Получение секретного ключа из переменной окружения
func GetSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

// Регистрация пользователя
func RegisterUser(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Сохранение в базу данных с контекстом
		user.Password = string(hashedPassword)  // если нужно, чтобы пароль был хеширован перед сохранением
		err = db.CreateUser(r.Context(), &user) // передаем контекст и указатель на структуру user

		if err != nil {
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// Логин пользователя
func LoginUser(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Проверка существования пользователя с контекстом
		storedUser, err := db.GetUserByEmail(context.Background(), user.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Проверка пароля
		err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Создание JWT токена
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    storedUser.ID,
			"email": storedUser.Email,
		})

		// Получение секретного ключа из переменной окружения
		secretKey := GetSecretKey()

		// Создание подписанного токена
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

// Middleware для проверки токена
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Проверка формата токена, должен быть "Bearer <token>"
		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[7:] // Обрезаем "Bearer "

		// Разбор и проверка токена
		secretKey := GetSecretKey()
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
