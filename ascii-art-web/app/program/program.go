package asciiartweb

import (
	"fmt"
	"net/http"

	routes "asciiartweb/app/routes"
)

var gConfigs *gconfigs = &gconfigs{HostURL: "localhost:8080"}

// Program - Execute Program With Default Json Configs
func Program() {
	SetJsonConfigs()
	ListenRoutes()
}

// ListenRoutes - Starting Listening All Routes in routes
func ListenRoutes() {
	ShowSuccesMessage(fmt.Sprintf("Start Listening Server: %v", gConfigs.HostURL))
	//Route For CSS
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("app/templates/css/"))))
	// Add Routes
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/ascii-art/", routes.Ascii_artHandler)
	// Start Listen
	err := http.ListenAndServe(gConfigs.HostURL, nil)
	if err != nil {
		ShowErrorMessage(err.Error())
	}
}
