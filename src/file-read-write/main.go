package main

import (
	"fmt"
	"io/ioutil"
)

//FileDs represents a file data structure
type FileDs struct {
	Name string
	Body []byte
}

//write body to a file
func (f *FileDs) write() error {
	fileName := f.Name + ".txt"
	return ioutil.WriteFile(fileName, f.Body, 0750)
}

//read a file in FileDs
func read(name string) (*FileDs, error) {
	fileName := name + ".txt"
	filebody, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return &FileDs{Name: name, Body: filebody}, nil
}

func main() {
	fileDs := &FileDs{Name: "test", Body: []byte("This is file content")}
	fileDs.write()

	outFile, err := read("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outFile.Body))
}
