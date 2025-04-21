package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func DummyLogin(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func TestDummyLogin(t *testing.T) {
	req, err := http.NewRequest("POST", "/dummyLogin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DummyLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}
