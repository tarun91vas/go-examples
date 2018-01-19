package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Case 1: Using built in Handler FileServer
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Case 2: Serving a subfolder inside public
	http.Handle("/text/", http.FileServer(http.Dir("public")))

	// Case 3: Serving a folder with different url path
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("public"))))

	fmt.Println("Server starting on 8080.....")
	http.ListenAndServe(":8080", nil)

	/* PS:
	Above code (Case 1) can be replaced by following line as well:
	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
	*/
}
