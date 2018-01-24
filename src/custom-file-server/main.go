package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//A custom file server: copies file content to http.ResponseWriter
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//set content type
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)

		io.Copy(w, f)
	})

	fmt.Println("Server starting on 8080.....")
	http.ListenAndServe(":8080", nil)
}
