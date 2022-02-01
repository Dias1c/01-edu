package handler

import (
	"forum/architecture/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

//
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

//
func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	// HERE IS ALL ROUTES
	return mux
}
