package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, found := pathToUrls[path]; found {
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

type pathUrl struct {
	Path string
	Url  string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	pathToUrls := buildMap(parsedYaml)
	return MapHandler(pathToUrls, fallback), nil
}

func parseYaml(ymlBytes []byte) ([]pathUrl, error) {
	parsed := []pathUrl{}
	err := yaml.Unmarshal(ymlBytes, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func buildMap(data []pathUrl) map[string]string {
	pathToUrls := map[string]string{}
	for _, obj := range data {
		pathToUrls[obj.Path] = obj.Url
	}

	return pathToUrls
}
