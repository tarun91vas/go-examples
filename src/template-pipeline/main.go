package main

import "html/template"
import "os"

//Person represents a person struct
type Person struct {
	Name string
	Age  int
	Rank int
}

//Sfactor defines some factor
func (p Person) Sfactor() float32 {
	return float32(p.Rank * 100 / p.Age)
}

const templateStr = ` 
{{- "Person Information"}}
Name: {{.Name}}
Age: {{.Age}}
Rank: {{.Rank}}
Sfactor: {{.Sfactor}}
`

func main() {
	p := Person{
		Name: "Goku",
		Age:  21,
		Rank: 5,
	}

	t := template.New("test")
	template.Must(t.Parse(templateStr))
	t.Execute(os.Stdout, p)
}
