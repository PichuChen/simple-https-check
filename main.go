package main

import (
	"net/http"

	"github.com/pichuchen/simple-https-check/web"
)

func main() {
	// make a http server and return "Hello World"
	http.HandleFunc("/", web.Handler)
	http.ListenAndServe(":8080", nil)

}
