package main

import (
	"fmt"
	"net/http"
)

// A common middleware function
func loggingMiddleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request url: " + r.URL.Path)
		hf(w, r)
	}
}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am first bruh!!"))
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Well, I am second!!"))
}

func main() {
	http.HandleFunc("/first", loggingMiddleware(firstHandler))

	http.HandleFunc("/second", loggingMiddleware(secondHandler))

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
