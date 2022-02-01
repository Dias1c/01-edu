package main

import (
	"forum/architecture/repository"
	"forum/architecture/service"
	"forum/architecture/web/handler"
	"forum/architecture/web/server"
	"log"
	"os"
)

func main() {
	repos := repository.NewRepo(nil)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(server.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
