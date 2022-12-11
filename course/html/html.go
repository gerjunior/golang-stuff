package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Title(urls ...string) <-chan string {
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
