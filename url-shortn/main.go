package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	urlshort "github.com/gerjunior/golang-stuff/url-shortn/urlshort"
)

func main() {
	mux := defaultMux()

	yamlFilePath := flag.String("yaml", "", "path to yaml file containing the mapped paths")
	jsonFilePath := flag.String("json", "", "path to json file containing the mapped paths")
	flag.Parse()

	yaml, err := os.ReadFile(*yamlFilePath)
	if err != nil {
		panic(err)
	}

	json, err := os.ReadFile(*jsonFilePath)
	if err != nil {
		panic(err)
	}

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler(json, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
