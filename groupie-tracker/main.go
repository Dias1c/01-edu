package main

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/routes"
	"log"
	"net/http"
	"os"
	"time"
)

const refreshAPIEveryMin = 5

func main() {
	ch := make(chan bool)
	go getData(ch)
	for success := range ch {
		if success {
			runServer()
		} else {
			log.Print("Can't run server without requested data")
			os.Exit(0)
		}
	}
}

func getData(ch chan bool) {
	t := time.Now()
	success := api.GetCache()
	if !success {
		ch <- false
		close(ch)
		return
	}
	api.GetSuggestions()
	log.Printf("Data parsed in %v", time.Since(t))
	ch <- true
	close(ch)
	for {
		time.Sleep(refreshAPIEveryMin * time.Minute)
		t = time.Now()
		api.RefreshAll()
		log.Printf("New data parsed in %v", time.Since(t))
	}
}

func runServer() {
	routes.InitTemplates()
	mux := http.NewServeMux()
	// Fs Assets
	assets := http.FileServer(http.Dir("public/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))
	// Front Pages
	mux.HandleFunc("/", routes.IndexHandler) // Index Page
	// API Pages
	mux.HandleFunc("/api/", routes.ArtistsHandler)                 // How Use API
	mux.HandleFunc("/api/artists/", routes.ArtistsHandler)         // All Artists
	mux.HandleFunc("/api/artist/", routes.ArtistsHandler)          // Artist By Id
	mux.HandleFunc("/api/filter/", routes.FilterHandler)           //Api Filter Page
	mux.HandleFunc("/api/suggestions/", routes.SuggestionsHandler) //Api Suggestions Page
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := fmt.Sprintf("localhost:%v", port)
	// Start Listen
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("Server started on http://%v", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err.Error())
	}
}
