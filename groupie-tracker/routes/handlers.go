package routes

import (
	"encoding/json"
	"groupie-tracker/api"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call page: %v\tHandleFunc: IndexHandler", r.URL.Path)
	if r.URL.Path != "/" && r.URL.Path != "/index" {
		renderError(w, http.StatusNotFound)
		return
	}
	// Check method
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, "index", nil)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	if r.Method == http.MethodOptions {
		return
	}
	log.Printf("Call page: %v\tHandleFunc: ArtistsHandler", r.URL.Path)
	// Check path
	if r.URL.Path != "/api/artists/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check method
	switch r.Method {
	case http.MethodGet:
		// Set the header and send data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*api.AllArtists)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	if r.Method == http.MethodOptions {
		return
	}
	log.Printf("Call page: %v\tHandleFunc: ArtistsHandler", r.URL.Path)
	// Check path
	if r.URL.Path != "/api/artist/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check method
	switch r.Method {
	case http.MethodGet:
		// Set the header and send data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*api.AllArtists)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	log.Printf("Call page: %v\tHandleFunc: SuggestionsHandler", r.URL.Path)
	if r.URL.Path != "/api/suggestions/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check method
	switch r.Method {
	case http.MethodGet:
		// Set the header and send data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*api.AllSuggestions)
	case http.MethodOptions:
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	log.Printf("Call page: %v\tHandleFunc: FilterHandler", r.URL.Path)
	if r.URL.Path != "/api/filter/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check header media
	value := r.Header.Get("Content-Type")
	if value != "" {
		if value != "application/json" {
			http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
			return
		}
	}
	// Check method
	switch r.Method {
	case http.MethodPost:
		// Decode the JSON from request body
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var filter api.FilterParams
		err := dec.Decode(&filter)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Check filter params
		success := api.CheckFilterFields(filter)
		if !success {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Filter with params
		result := api.FilterByParams(filter)
		// Set the header and send data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*result)
	case http.MethodOptions:
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
