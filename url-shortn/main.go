package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gerjunior/golang-stuff/url-shortn/db"
	urlshort "github.com/gerjunior/golang-stuff/url-shortn/urlshort"
)

var ctx = context.Background()

// docker run -p 6379:6379 --name some-redis -d redis
// go run main.go -yaml ./content/paths.yml -json ./content/paths.json
func main() {
	db.RedisInit()
	err := db.Rdb.Set(ctx, "/redis", "https://hub.docker.com/_/redis", time.Hour).Err()
	if err != nil {
		panic(err)
	}
	db.Rdb.Set(ctx, "/youtube", "https://www.youtube.com/watch?v=1WVcZg9BWSM", time.Hour).Err()
	if err != nil {
		panic(err)
	}
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
	redisHandler := urlshort.RedisHandler(ctx, jsonHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", redisHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
