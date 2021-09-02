package routes

import "net/http"

type errorParams struct {
	Message string
}

func renderError(w http.ResponseWriter, code int) {
	// http.Error(w, msg, code)
	w.WriteHeader(code)
	err := &errorParams{Message: http.StatusText(code)}
	renderTemplate(w, "error", *err)
}
