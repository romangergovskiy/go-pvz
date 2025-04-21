package pvz

import (
	"encoding/json"
	"net/http"

	"github.com/romangergovskiy/go-pvz/internal/database"
)

func CreatePVZ(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "PVZ created"})
	}
}

func GetPVZ(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "PVZ fetched"})
	}
}
