package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "To get in touch, please contact our support at <a href=\"mailto:support@mygallery.in\">support@mygallery.in</a>.")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}
