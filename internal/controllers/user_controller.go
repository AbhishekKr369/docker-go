package controllers

import (
	"encoding/json"
	"go-medium-project/internal/models"

	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{Id: 1, Name: "Abhi"},
		{Id: 2, Name: "Divya"},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("New User Created")
}
