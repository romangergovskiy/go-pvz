package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type User struct {
	ID       int
	Email    string
	Password string
}

// Инициализация базы данных
func InitDB() (*DB, error) {
	connStr := "user=postgres dbname=pvz password=yourpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Создание пользователя в базе данных
func (db *DB) CreateUser(email, password string) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := db.Exec(query, email, password)
	return err
}

// Получение пользователя по email
func (db *DB) GetUserByEmail(email string) (User, error) {
	var user User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
