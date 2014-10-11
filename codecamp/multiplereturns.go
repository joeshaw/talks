package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("http://google.com")
	fmt.Println(err)

	_, err = http.Get("ftp://google.com")
	fmt.Println(err)
}
