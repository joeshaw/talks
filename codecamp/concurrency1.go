package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// START OMIT
func define(word string) string {
	out, err := exec.Command("/usr/bin/python", "dict.py", word).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func main() {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if strings.Contains(s.Text(), "code") {
			fmt.Println(define(s.Text()))
		}
	}
}

// END OMIT
