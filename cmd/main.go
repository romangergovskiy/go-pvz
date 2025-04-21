package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/romangergovskiy/go-pvz/internal/auth"
	"github.com/romangergovskiy/go-pvz/internal/database"
	"github.com/romangergovskiy/go-pvz/internal/pvz"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env файл не найден или не загружен, используем переменные окружения по умолчанию")
	} else {
		log.Println("✅ .env файл успешно загружен")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/register", auth.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/login", auth.LoginUser(db)).Methods("POST")

	secured := router.PathPrefix("/secured").Subrouter()
	secured.Use(auth.VerifyToken)

	secured.HandleFunc("/pvz", pvz.CreatePVZ(db)).Methods("POST")
	secured.HandleFunc("/pvz/{id}", pvz.GetPVZ(db)).Methods("GET")

	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
