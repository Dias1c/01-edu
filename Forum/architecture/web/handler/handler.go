package handler

import (
	"fmt"
	"forum/architecture/service"
	"net/http"
	"path/filepath"
	"text/template"
)

type Handler struct {
	service   *service.Service
	templates *template.Template
}

//
func NewHandler(service *service.Service) (*Handler, error) {
	templates, err := GetTemplate()
	if err != nil {
		return nil, fmt.Errorf("NewHandler: %w", err)
	}
	return &Handler{
		service:   service,
		templates: templates,
	}, nil
}

//
func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	// HERE IS ALL ROUTES
	fsStatic := http.FileServer(http.Dir("templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	mux.HandleFunc("/test", h.TestHandler)
	return mux
}

//
func GetTemplate() (*template.Template, error) {
	// Gets All Templates in folder templates
	filepaths, err := filepath.Glob("templates/*.gohtml")
	files, err := template.ParseFiles(filepaths...)
	if err != nil {
		return nil, fmt.Errorf("GetTemplate: %w", err)
	}
	return template.Must(files, nil), nil
}
