package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func main() {
	file, err := os.ReadFile("./story.json")
	if err != nil {
		panic(err)
	}

	book := Book{}
	err = json.Unmarshal(file, &book)
	if err != nil {
		panic(err)
	}

	arc := book["intro"]

	for {
		fmt.Printf("%s\n\n\n", arc.Title)

		for _, paragraph := range arc.Story {
			fmt.Printf("%s\n\n", paragraph)
		}

		if len(arc.Options) == 0 {
			fmt.Println("The End")
			break
		}

		fmt.Println("Choose an option:")

		for idx, option := range arc.Options {
			fmt.Printf("#%d) %s\n", idx+1, option.Text)
		}

		answerCh := make(chan int)

		go func() {
			for {
				var input string
				fmt.Scanf("%s", &input)
				parsed, err := strconv.Atoi(input)
				if err != nil {
					continue
				}

				if parsed > len(arc.Options) {
					continue
				}

				answerCh <- parsed
				break
			}
		}()

		select {
		case answer := <-answerCh:
			chosenArc := arc.Options[answer-1]
			arc = book[chosenArc.Arc]
		}
	}
}
