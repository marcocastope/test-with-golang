package models

import (
	"encoding/json"
	"net/http"
)

// User ...
type User struct {
	UserName string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	data := []User{
		User{
			UserName: "marcocastope",
			Fullname: "Marco Luis Castope",
			Email:    "marcocastope7@gmail.com",
		},
		User{
			UserName: "elángeldemadrid",
			Fullname: "Raul Gonzales Blanco",
			Email:    "raulgonzales@gmail.com",
		},
		User{
			UserName: "elmagodelbalón",
			Fullname: "Zinedine Zidane",
			Email:    "zz@gmail.com",
		},
	}

	j, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
