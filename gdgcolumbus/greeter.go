package main

import (
	"fmt"
	"strings"
)

type Greeter interface {
	Greet() string
}

type Dog struct {
	Name    string
	Breed   string
	GoodBoy bool
}

func (d Dog) Greet() string {
	return "Woof!"
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, %s.", p.Name)
}

// START OMIT
func greetLoudly(g Greeter) string {
	greeting := g.Greet()
	return strings.ToUpper(greeting)
}

func main() {
	p := Person{Name: "Jessica"}
	fmt.Println(greetLoudly(p))

	d := Dog{Name: "Wolfy"}
	fmt.Println(greetLoudly(d))
}

// END OMIT
