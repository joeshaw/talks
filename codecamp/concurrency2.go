package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func define(word string) string {
	out, err := exec.Command("/usr/bin/python", "dict.py", word).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

// START PAGE 1 OMIT
func main() {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	in := make(chan string)
	out := make(chan string)

	// Machine 1: Read words containing "code"
	go func() {
		for s.Scan() {
			if strings.Contains(s.Text(), "code") {
				in <- s.Text()
			}
		}
		close(in)
	}()
	// END PAGE 1 OMIT

	// START PAGE 2 OMIT
	// Machine 2: Gather definitions of words in parallel
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for word := range in {
				out <- define(word)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// Machine 3: Print out definitions
	for def := range out {
		fmt.Println(def)
	}
}

// END PAGE 2 OMIT
