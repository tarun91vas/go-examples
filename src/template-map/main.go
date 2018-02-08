package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Parse all the templates at a path
	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[1:]
		t := templates.Lookup(filePath + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	fmt.Println("Server starting on 8080....")
	http.ListenAndServe(":8080", nil)

}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basepath = "templates"
	template.Must(result.ParseGlob(basepath + "/*.html"))
	return result
}
