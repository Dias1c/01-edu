package handler

import (
	"log"
	"net/http"
)

// Handler For Test
func (h *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "pg-index", nil)
	err = h.templates.ExecuteTemplate(w, "pg-login", nil)
	err = h.templates.ExecuteTemplate(w, "pg-question", nil)
	err = h.templates.ExecuteTemplate(w, "pg-questions", nil)
	err = h.templates.ExecuteTemplate(w, "pg-signup", nil)
	err = h.templates.ExecuteTemplate(w, "pg-login", nil)
	err = h.templates.ExecuteTemplate(w, "pg-tags", nil)
	err = h.templates.ExecuteTemplate(w, "pg-user", nil)
	err = h.templates.ExecuteTemplate(w, "pg-users", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
