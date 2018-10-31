package main

import (
	"fmt"
	"net/http"
)

// still making two calls to server ?..TBD
func main() {
	// Using built in http function ServeFile
	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		// Check to see if writer implements pusher interface
		if pusher, ok := w.(http.Pusher); ok {
			fmt.Println("writer implements pusher interface....")
			err := pusher.Push("/css/main.css", &http.PushOptions{
				Header: http.Header{"Content-Type": []string{"text/css"}},
			})

			if err != nil {
				fmt.Println("Oops! Server push failed", err)
			}
		}
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("Server starting on 8080.....")
	http.ListenAndServeTLS(":8080", "dummy-certs/cert.pem", "dummy-certs/key.pem", nil)
}
