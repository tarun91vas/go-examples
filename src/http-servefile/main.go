package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Using built in http function ServeFile
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public")
	})
	fmt.Println("Server starting on 8080.....")
	http.ListenAndServe(":8080", nil)
}
