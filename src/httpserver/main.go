package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http server using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello hooman!! Server responded with plain text string!!"))
	})

	fmt.Println("Listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
