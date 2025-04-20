package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/romangergovskiy/go-pvz/internal/auth"
	"github.com/romangergovskiy/go-pvz/internal/database"
	"github.com/romangergovskiy/go-pvz/internal/pvz"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	// Инициализация маршрутизатора
	router := mux.NewRouter()

	// Маршруты для регистрации и логина
	router.HandleFunc("/register", auth.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/login", auth.LoginUser(db)).Methods("POST")

	// Защищённые маршруты
	secured := router.PathPrefix("/secured").Subrouter()
	secured.Use(auth.VerifyToken)
	secured.HandleFunc("/pvz", pvz.CreatePVZ(db)).Methods("POST")
	secured.HandleFunc("/pvz/{id}", pvz.GetPVZ(db)).Methods("GET")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8000", router))
}
