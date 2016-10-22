// BEGIN 1 OMIT

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

// END 1 OMIT
// BEGIN 2 OMIT

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

	// END 2 OMIT
	// BEGIN 3 OMIT

	s := bufio.NewScanner(in)
	for s.Scan() {
		hostname := s.Text()

		fastly := isFastly(hostname)
		if !fastly && !strings.HasPrefix(hostname, "www.") {
			fastly = isFastly("www." + hostname)
		}

		if fastly {
			fmt.Fprintln(mw, hostname)
		}
	}
}

// END 3 OMIT
