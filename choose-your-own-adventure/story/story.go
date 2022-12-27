package story

import (
	"encoding/json"
	"os"
)

type StoryOptions struct {
	Text string
	Arc  string
}

type StoryArc struct {
	Title   string
	Story   []string
	Options []StoryOptions
}

type Book map[string]StoryArc

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
