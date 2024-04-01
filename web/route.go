package web

import (
	"net/http"

	"github.com/pichuchen/simple-https-check/web/root"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", root.HandleFunc)

	mux.ServeHTTP(w, r)

}
