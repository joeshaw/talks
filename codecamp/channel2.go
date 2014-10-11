package main

import (
	"fmt"
	"net/http"
	"time"
)

func getURL(ch chan string, url string) {
	start := time.Now()
	http.Get(url)
	ch <- fmt.Sprintf("Fetched %s in %s", url, time.Since(start))
}

// START OMIT
func main() {
	ch := make(chan string)

	go getURL(ch, "http://google.com")
	go getURL(ch, "http://yahoo.com")
	go getURL(ch, "http://bing.com")
	go getURL(ch, "http://duckduckgo.com")

	timeout := time.After(500 * time.Millisecond)

	for {
		select {
		case s := <-ch:
			fmt.Println(s)

		case <-timeout:
			fmt.Println("Timeout expired!")
			return
		}
	}
}

// END OMIT
