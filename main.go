package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "To get in touch, please contact our support at <a href=\"mailto:support@mygallery.in\">support@mygallery.in</a>.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8000", r)
}
