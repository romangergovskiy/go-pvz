package models

type User struct {
	ID           int
	Email        string
	PasswordHash string // Если ты хранишь захешированный пароль
	Password     string // Если ты хранишь оригинальный пароль (не рекомендуется)
}
