package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http server using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using TLS for communication"))
	})

	fmt.Println("Listening on 8080...")
	// Pass certs for authorization
	http.ListenAndServeTLS(":8080", "dummy-certs/cert.pem", "dummy-certs/key.pem", nil)
}
