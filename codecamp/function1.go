package main

import "fmt"

// START OMIT
func repeat(n int, fn func()) {
	for i := 0; i < n; i++ {
		fn()
	}
}

func helloGophers() {
	fmt.Println("Hello Gophers!")
}

func main() {
	repeat(3, helloGophers)
}

// END OMIT
