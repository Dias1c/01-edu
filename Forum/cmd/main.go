package main

import (
	"forum"
	"forum/repository"
	"forum/service"
	"forum/web/api"
	handler "forum/web/api"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/posts", api.GetAllPosts)

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(forum.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
