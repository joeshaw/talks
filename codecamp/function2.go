package main

import "fmt"

func repeat(n int, fn func()) {
	for i := 0; i < n; i++ {
		fn()
	}
}

// START OMIT
func main() {
	start := 2

	repeat(10, func() {
		fmt.Printf("Powers of two: %d\n", start)
		start *= 2
	})
}

// END OMIT
