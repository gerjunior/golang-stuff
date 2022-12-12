package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gerjunior/golang-stuff/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("server started at :9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
