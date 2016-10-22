package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func isFastly(hostname string) bool {
	cname, _ := net.LookupCNAME(hostname)
	return strings.Contains(cname, "fastly")
}

// BEGIN 1 OMIT
func readHostnames(in io.Reader, inCh chan<- string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		inCh <- s.Text() // HL
	}
}

// END 1 OMIT

// BEGIN 2 OMIT
func lookupHostnames(inCh <-chan string, outCh chan<- string) {
	for hostname := range inCh { // HL
		fastly := isFastly(hostname)
		if !fastly && !strings.HasPrefix(hostname, "www.") {
			fastly = isFastly("www." + hostname)
		}

		if fastly {
			outCh <- hostname // HL
		}
	}
}

// END 2 OMIT

// BEGIN 3 OMIT
func outputFastlyHostnames(out io.Writer, outCh <-chan string) {
	for hostname := range outCh { // HL
		fmt.Fprintln(out, hostname)
	}
}

// END 3 OMIT

func main() {
	in, err := os.Open("1h.txt")
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

	// BEGIN 4 OMIT
	inCh := make(chan string)
	outCh := make(chan string)

	go func() {
		readHostnames(in, inCh)
		close(inCh)
	}()

	go func() {
		lookupHostnames(inCh, outCh)
		close(outCh)
	}()

	outputFastlyHostnames(mw, outCh)
	// END 4 OMIT
}
