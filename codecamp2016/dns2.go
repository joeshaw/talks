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

	// BEGIN OMIT
	s := bufio.NewScanner(in)
	for s.Scan() {
		hostname := s.Text()

		go func() {
			fastly := isFastly(hostname)
			if !fastly && !strings.HasPrefix(hostname, "www.") {
				fastly = isFastly("www." + hostname)
			}

			if fastly {
				fmt.Fprintln(mw, hostname)
			}
		}()
	}
	// END OMIT
}
