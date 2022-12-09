package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed")
		w.Write([]byte("Request processed"))
	case <-ctx.Done():
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
