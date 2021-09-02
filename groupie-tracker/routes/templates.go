package routes

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

func InitTemplates() {
	files, err := template.ParseFiles("public/index.html", "public/error.html")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates = template.Must(files, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, vars interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", vars)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
