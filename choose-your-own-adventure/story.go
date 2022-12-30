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
	b      Book
	t      *template.Template
	pathFn func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := (h.b)[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Something went wrong...", http.StatusBadRequest)
		}
		return
	}

	http.Error(w, "Chapter not found...", http.StatusBadRequest)
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(b Book, opts ...HandlerOption) http.Handler {
	defaultTmpl := template.Must(template.New("Template.html").ParseFiles("Template.html"))
	h := handler{b, defaultTmpl, pathFn}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}
