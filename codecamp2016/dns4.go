package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func isFastly(hostname string) bool {
retry:
	cname, err := net.LookupCNAME(hostname)
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return false
		}
		if err, ok := err.(*net.DNSError); ok && (err.Timeout() || err.Temporary()) {
			fmt.Println("RETRY", hostname, err)
			goto retry
		}
		fmt.Println("ERROR", hostname, err)
		return false
	}
	return strings.Contains(cname, "fastly")
}

func readHostnames(in io.Reader, inCh chan<- string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		inCh <- s.Text()
	}
}

func lookupHostnames(inCh <-chan string, outCh chan<- string) {
	for hostname := range inCh {
		fastly := isFastly(hostname)
		if !fastly && !strings.HasPrefix(hostname, "www.") {
			fastly = isFastly("www." + hostname)
		}

		if fastly {
			outCh <- hostname
		}
	}
}

func outputFastlyHostnames(out io.Writer, outCh <-chan string) {
	for hostname := range outCh {
		fmt.Fprintln(out, hostname)
	}
}

func main() {
	in, err := os.Open("1m.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create("fastly.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	mw := io.MultiWriter(os.Stdout, out)

	inCh := make(chan string)
	outCh := make(chan string)

	// BEGIN OMIT
	const nWorkers = 200

	go func() {
		readHostnames(in, inCh)
		close(inCh)
	}()

	var wg sync.WaitGroup
	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			lookupHostnames(inCh, outCh)
			wg.Done()
		}()
	}

	go func() {
		// Wait until all lookupHostnames() exit
		wg.Wait()
		close(outCh)
	}()

	outputFastlyHostnames(mw, outCh)
	// END OMIT
}
