package api

import (
	"forum/architecture/service"
	"net/http"
)

//
type ApiHandler struct {
	service *service.Service
}

//
func NewApiHandler(service *service.Service) *ApiHandler {
	return &ApiHandler{
		service: service,
	}
}

//
func (v *ApiHandler) InitRoutes(mux *http.ServeMux) {
	// Init Api routes
}
