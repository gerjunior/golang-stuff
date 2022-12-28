package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type ChapterOptions struct {
	Text    string
	Chapter string `json:"arc"`
}

type BookChapter struct {
	Title      string
	Paragraphs []string `json:"story"`
	Options    []ChapterOptions
}

type Book map[string]BookChapter

func ParseBook(bookPath string) (Book, error) {
	file, err := os.Open(bookPath)
	if err != nil {
		return nil, err
	}

	book := Book{}
	err = json.NewDecoder(file).Decode(&book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

type handler struct {
	b Book
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	if chapter, ok := (h.b)[path[1:]]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Something went wrong...", http.StatusBadRequest)
		}
		return
	}

	http.Error(w, "Chapter not found...", http.StatusBadRequest)
}

func NewHandler(b Book, t *template.Template) http.Handler {
	return handler{b, t}
}
