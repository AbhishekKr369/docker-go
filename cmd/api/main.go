package main

import (
	"go-medium-project/internal/config"
	"go-medium-project/internal/routers"
	"net/http"
)

func main() {
	config := config.LoadConfig()
	r := routers.SetUpRouter()
	http.ListenAndServe(config.ServerPort, r)

}
