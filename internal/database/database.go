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
	Conn *sql.DB
}

func InitDB() (*DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{Conn: db}, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}

func (db *DB) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id`
	return db.Conn.QueryRowContext(ctx, query, user.Email, user.Password, user.Role).Scan(&user.ID)
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password, role FROM users WHERE email = $1`
	row := db.Conn.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
