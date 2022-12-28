package story

import (
	"encoding/json"
	"os"
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
