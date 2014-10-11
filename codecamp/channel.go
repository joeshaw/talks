package main

import (
	"fmt"
	"net/http"
	"time"
)

// START OMIT
func getURL(ch chan string, url string) {
	start := time.Now()
	http.Get(url)
	ch <- fmt.Sprintf("Fetched %s in %s", url, time.Since(start))
}

func main() {
	start := time.Now()
	ch := make(chan string)

	go getURL(ch, "http://google.com")
	go getURL(ch, "http://yahoo.com")
	go getURL(ch, "http://bing.com")
	go getURL(ch, "http://duckduckgo.com")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("Total runtime: %s\n", time.Since(start))
}

// END OMIT
