// START 1 OMIT
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// END 1 OMIT

// START 2 OMIT
func fetch(site string) string {
	start := time.Now()

	resp, err := http.Get("http://" + site)
	if err != nil {
		return fmt.Sprintf("error fetching %s: %s", site, err)
	}
	defer resp.Body.Close()

	f, err := os.Create(fmt.Sprintf("/tmp/%s", site))
	if err != nil {
		return fmt.Sprintf("unable to save %s to %s: %s", site, f.Name(), err)
	}
	defer f.Close()

	io.Copy(f, resp.Body) // Ignoring errors for brevity.  Never do that.

	return fmt.Sprintf("fetched %s in %s", site, time.Since(start))
}

// END 2 OMIT

// START 3 OMIT
func main() {
	resultCh := make(chan string)
	sites := []string{"google.com", "yahoo.com", "bing.com", "duckduckgo.com", "joeshaw.org"}

	for i := range sites {
		site := sites[i]
		go func() {
			resultCh <- fetch(site)
		}()
	}

	for res := range resultCh {
		fmt.Println(res)
	}
}

// END 3 OMIT
