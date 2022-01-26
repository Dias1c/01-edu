package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var Templates *template.Template

func main() {
	var err error
	err = InitTemplates()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Init Handlers + Run Server
	mux := http.NewServeMux()
	// FS
	fs := http.FileServer(http.Dir("templates/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// Pages
	mux.HandleFunc("/", IndexHandler)        // Index Page
	mux.HandleFunc("/create", CreateHandler) // Create Page
	port := ":8080"
	server := http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("Server started on http://localhost%v", port)
	err = server.ListenAndServe()
}

// InitTemplates - init	Templates
func InitTemplates() error {
	filenames := []string{
		"templates/index.gohtml",
		"templates/header.gohtml",
		"templates/footer.gohtml",
		"templates/question-create.gohtml",
	}
	files, err := template.ParseFiles(filenames...)
	if err != nil {
		return err
	}
	Templates = template.Must(files, nil)
	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "question-create.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
