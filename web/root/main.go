package root

import (
	"net/http"
	"os"

	_ "embed"
)

//go:embed template.html
var templateFile []byte

// GET /
func HandleFunc(w http.ResponseWriter, r *http.Request) {

	newTemplate, err := os.ReadFile("web/root/template.html")
	if err == nil {
		// if the file is found, renew the template
		templateFile = newTemplate
	}

	w.Write(templateFile)

}
