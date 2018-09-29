package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/marcocastope/TEST-WITH-GOLANG/models"
)

func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", models.GetUsers).Methods("GET")

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("el código de estado esperado es: 200, y tiene un código de estado actual de %d", w.Code)
	}

}
