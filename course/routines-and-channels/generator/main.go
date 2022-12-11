package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	c := make(chan string, 2)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			matches := r.FindStringSubmatch(string(html))
			if len(matches) == 0 {
				c <- "Error: " + url
			} else {
				c <- matches[1]
			}
		}(url)
	}

	return c
}

func main() {
	t1 := title("https://www.google.com", "https://www.cod3r.com.br")
	t2 := title("https://facebook.com", "https://www.teste.com")
	fmt.Println("First: ", <-t1, " | ", <-t2)
	fmt.Println("Second: ", <-t1, " | ", <-t2)
}
