package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func forward(from <-chan string, to chan string) {
	for {
		to <- <-from
	}
}

// Multiplex - Mix string streaming into chan
func aggregate(in1, in2 <-chan string) <-chan string {
	ch := make(chan string)
	go forward(in1, ch)
	go forward(in2, ch)
	return ch
}

// Generator
func title(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}

func main() {
	ch := aggregate(
		title("https://www.cod3r.com.br", "https://www.google.com"),
		title("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
