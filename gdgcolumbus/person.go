package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, %s.", p.Name)
}

func main() {
	p := Person{
		Name: "Jessica",
		Age:  44,
	}

	fmt.Println(p.Greet())
}
