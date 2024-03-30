package main

import (
	"fmt"
	"net/http"
)

func main() {
	// make a http server and return "Hello World"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.ListenAndServe(":8080", nil)

}
