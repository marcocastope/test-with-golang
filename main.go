package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcocastope/TEST-WITH-GOLANG/models"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", models.GetUsers).Methods("GET")
	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
