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
	chapter := strings.Split(r.URL.Path, "/")[1]
	content := (*book)[chapter]

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("There was an error while executing the template for this book."))
	}
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
