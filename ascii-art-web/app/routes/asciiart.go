package asciiartweb

import (
	"fmt"
	"net/http"
)

func Ascii_artHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "ASCII", "Handler")
}
