package main

import (
	"net/http"

	"github.com/Le0nar/game_shop/internal/handler"
	"github.com/Le0nar/game_shop/internal/repository"
	"github.com/Le0nar/game_shop/internal/service"
)

func main() {
	// TODO: init and throw db to NewRepository
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}
