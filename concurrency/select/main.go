package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

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

func faster(url1, url2, url3 string) string {
	ch1 := title(url1)
	ch2 := title(url2)
	ch3 := title(url3)

	select {
	case t1 := <-ch1:
		return t1
	case t2 := <-ch2:
		return t2
	case t3 := <-ch3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "All lost"
	}
}

func main() {
	fmt.Println(faster(
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	))
}
