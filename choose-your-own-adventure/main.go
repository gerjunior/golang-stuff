package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gerjunior/golang-stuff/choose-your-own-adventure/story"
)

func main() {
	port := flag.Int("p", 8080, "port to run the web app")
	filePath := flag.String("f", "./story.json", "book json file path")
	flag.Parse()

	parsedBook, err := story.ParseBook(*filePath)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("Template.html").ParseFiles("Template.html"))
	handler := story.NewHandler(parsedBook, tmpl)

	http.HandleFunc("/", handler.ServeHTTP)

	fmt.Printf("Listening on port :%d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
