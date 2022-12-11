package main

import (
	"fmt"
	"time"

	"github.com/gerjunior/golang-stuff/course/html"
)

func forwardFromTo(origin <-chan string, destiny chan string) {
	for {
		select {
		case item := <-origin:
			destiny <- item
		case <-time.After(time.Second):
			close(destiny)
		}
	}
}

func merge(channels ...<-chan string) <-chan string {
	merged := make(chan string)

	for _, channel := range channels {
		go forwardFromTo(channel, merged)
	}

	return merged
}

func main() {
	start := time.Now()
	c := merge(
		html.Title("https://google.com", "https://www.cod3r.com.br"),
		html.Title("https://facebook.com", "https://www.amazon.com.br"),
	)

	for item := range c {
		fmt.Println(item)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
