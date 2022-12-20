package urlshort

import (
	"fmt"
	"net/http"
)

func registerHandleFunc(short, path string, fallback http.Handler) {
	if path == "" {
		http.Handle(short, fallback)
		return
	}

	fmt.Printf("Mapping %s to %s\n", short, path)
	http.HandleFunc(short, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, http.StatusSeeOther)
	})
}

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range pathToUrls {
			registerHandleFunc(key, value, fallback)
		}
	})
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
