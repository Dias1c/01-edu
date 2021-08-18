package asciiartweb

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("app/templates/index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
