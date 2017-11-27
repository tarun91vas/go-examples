package main

import "net/http"
import "fmt"

type myHandler struct {
	name string
}

// myHandler implements ServeHTTP function for Handler
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %s", mh.name)))
}

func main() {
	// Implment ServeHTTP for custom handler
	http.Handle("/", &myHandler{name: "Gopher"})

	fmt.Printf("Listening on 8080....")
	http.ListenAndServe(":8080", nil)
}
