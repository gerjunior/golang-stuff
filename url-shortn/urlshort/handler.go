package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirect(w, r, pathToUrls, fallback)
	}
}

func redirect(w http.ResponseWriter, r *http.Request, pathToUrls map[string]string, fallback http.Handler) {
	path := r.URL.Path
	url := pathToUrls[path]

	if url != "" {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	fallback.ServeHTTP(w, r)
}

type URLs struct {
	Path string
	Url  string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsed := []URLs{}
	err := yaml.Unmarshal(yml, &parsed)

	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string, 0)
	for _, obj := range parsed {
		pathToUrls[obj.Path] = obj.Url
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		redirect(w, r, pathToUrls, fallback)
	}

	return handler, nil
}
