package order

import (
	"net/http"
)

type Order struct {
	ID     string `json:"id"`
	PVZID  string `json:"pvzId"`
	Status string `json:"status"`
}

func AcceptGoods(w http.ResponseWriter, r *http.Request) {
	// Логика начала приёма товаров
}

func AddGoodsToOrder(w http.ResponseWriter, r *http.Request) {
	// Логика добавления товаров в приёмку
}

func CloseOrder(w http.ResponseWriter, r *http.Request) {
	// Логика закрытия приёмки
}
