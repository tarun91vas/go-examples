package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Create a request handler for pattern `text`
	http.HandleFunc("/text/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/text/")
	})

	// Redirect root request to `text`` handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/text/", http.StatusFound)
	})

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
