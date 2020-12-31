package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	p1 := &Page{Title: "titulo", Body: []byte("contenido")}
	p1.save()
	//p2, _ := loadPage("titulo")
	//fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title+".txt", p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	body, error := ioutil.ReadFile(title+".txt")
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
