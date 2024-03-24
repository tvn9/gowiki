package main

// First building block of a web app

import (
	"fmt"
	"os"
)

// Page structure to hold a page Title and Body
type Page struct {
	Title string
	Body  []byte
}

// save function save the input text to a file
func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

// loadPage function read a file content and store the file data in Page struct
// fileName will be stored in Title field and file content will be stove in Body.
func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// main function start the program
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("Thanh Phố Hồ Chí Minh.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
