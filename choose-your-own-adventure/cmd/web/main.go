package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	story "github.com/gerjunior/golang-stuff/choose-your-own-adventure"
)

func main() {
	port := flag.Int("p", 8080, "port to run the web app")
	filePath := flag.String("f", "./story.json", "book json file path")
	flag.Parse()

	parsedBook, err := story.ParseBook(*filePath)
	if err != nil {
		panic(err)
	}

	// handler := story.NewHandler(parsedBook, story.WithTemplate(nil))
	handler := story.NewHandler(parsedBook)

	http.HandleFunc("/", handler.ServeHTTP)

	fmt.Printf("Listening on port :%d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
