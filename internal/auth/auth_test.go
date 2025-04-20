package auth

import (
	// Путь к вашему пакету database
	"net/http"
	"net/http/httptest"
	"testing"
)

// Пример функции для тестирования DummyLogin
func DummyLogin(w http.ResponseWriter, r *http.Request) {
	// Логика обработки запроса
	// Для примера, просто отвечаем с успешным статусом
	w.WriteHeader(http.StatusOK)
}

func TestDummyLogin(t *testing.T) {
	// Создаём HTTP запрос к /dummyLogin
	req, err := http.NewRequest("POST", "/dummyLogin", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаём новый Recorder для захвата ответа
	rr := httptest.NewRecorder()

	// Создаём обработчик
	handler := http.HandlerFunc(DummyLogin)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}
