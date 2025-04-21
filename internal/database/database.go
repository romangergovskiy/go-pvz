package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/romangergovskiy/go-pvz/internal/models"
)

type DB struct {
	*sql.DB
}

func InitDB() (*DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия подключения: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ошибка пинга базы данных: %v", err)
	}

	return &DB{db}, nil
}

func (db *DB) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err := db.ExecContext(ctx, query, user.Email, user.PasswordHash)
	return err
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	row := db.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя по email: %v", err)
	}

	return &user, nil
}
