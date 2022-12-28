package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gerjunior/golang-stuff/choose-your-own-adventure/story"
)

var book *story.Book
var tmpl *template.Template

func chapterHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	if chapter, ok := (*book)[path[1:]]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Something went wrong...", http.StatusBadRequest)
		}
		return
	}

	http.Error(w, "Chapter not found...", http.StatusBadRequest)
}

func main() {
	port := flag.Int("p", 8080, "port to run the web app")
	filePath := flag.String("f", "./story.json", "book json file path")
	flag.Parse()

	parsedBook, err := story.ParseBook(*filePath)
	if err != nil {
		panic(err)
	}

	book = &parsedBook
	tmpl = template.Must(template.New("Template.html").ParseFiles("Template.html"))

	http.HandleFunc("/", chapterHandler)

	fmt.Printf("Listening on port :%d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
