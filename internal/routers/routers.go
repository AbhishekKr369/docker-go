package routers

import (
	"go-medium-project/internal/controllers"

	"github.com/gorilla/mux"
)

func SetUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	return r
}
