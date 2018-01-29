package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	// string template
	templateString := `shi, I am a dumb template`

	// create and parse a new template
	t, err := template.New("Test").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}

	// execute template to standard output
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

}
