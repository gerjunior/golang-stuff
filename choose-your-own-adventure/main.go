package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gerjunior/golang-stuff/choose-your-own-adventure/story"
)

func read(book story.Book) {
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

func main() {
	filePath := flag.String("f", "./story.json", "book json file path")
	flag.Parse()

	book, err := story.ParseBook(*filePath)
	if err != nil {
		panic(err)
	}

	read(book)
}
