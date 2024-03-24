package main

// Create a simple web server to load the content in a text file as web content

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Biên Hoà Đồng Nai - Việt Nam! %s</h1>", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

const text = `
Thanh Pho Ho Chi Minh<br>
Đảng Cộng Sản Sống Mãi Trong Quần Chúng Ta.<br>
Như Có Bác Hồ Trong Ngày Vui Đại Thắng!
`

func main() {
	p := Page{Title: "TestPage", Body: []byte(text)}
	p.save()

	http.HandleFunc("/", home)
	http.HandleFunc("/view/", viewHandler)

	fmt.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
