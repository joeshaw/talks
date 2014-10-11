package main

type Foo struct{}

func main() {
	// START OMIT
	a := "hello world"
	a = 4 // compile-time error
	// END OMIT
}
