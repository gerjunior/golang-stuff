package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gerjunior/golang-stuff/choose-your-own-adventure/story"
)

func returnError(w http.ResponseWriter, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	filePath := flag.String("f", "./story.json", "book json file path")
	flag.Parse()

	book, err := story.ParseBook(*filePath)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("Template.html").ParseFiles("Template.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc := strings.Split(r.URL.Path, "/")[1]

		story := book[arc]

		err = tmpl.Execute(w, story)
		if err != nil {
			fmt.Println(err)
			returnError(w, map[string]interface{}{
				"error": "Unable to execute template",
			})
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
