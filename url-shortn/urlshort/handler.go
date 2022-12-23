package urlshort

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gerjunior/golang-stuff/url-shortn/db"
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

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJson(jsonBytes)
	if err != nil {
		return nil, err
	}

	pathToUrls := buildMap(parsedJson)
	return MapHandler(pathToUrls, fallback), nil
}

func parseJson(jsonBytes []byte) ([]pathUrl, error) {
	parsed := []pathUrl{}
	err := json.Unmarshal(jsonBytes, &parsed)
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

func RedisHandler(ctx context.Context, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		url, err := db.Rdb.Get(ctx, path).Result()
		if err != nil {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
