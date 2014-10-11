package main

import (
	"fmt"
	"math"
)

func main() {
	sin, cos := math.Sincos(math.Pi / 2.0)
	fmt.Printf("%f %f\n", sin, cos)
}
