package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/romangergovskiy/go-pvz/internal/database"
	"github.com/romangergovskiy/go-pvz/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func GetSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func RegisterUser(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)

		if err := db.CreateUser(r.Context(), &user); err != nil {
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User registered",
			"userID":  user.ID,
		})
	}
}

func LoginUser(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByEmail(r.Context(), creds.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
		})
		tokenString, err := token.SignedString([]byte(GetSecretKey()))
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
