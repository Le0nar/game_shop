package main

import (
	"net/http"

	"github.com/Le0nar/game_shop/internal/handler"
	"github.com/Le0nar/game_shop/internal/repository"
	"github.com/Le0nar/game_shop/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.NewPostgresDB(repository.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "gameshop",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}
