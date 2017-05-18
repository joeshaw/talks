package main

import "fmt"

type Dog struct {
	Name    string
	Breed   string
	GoodBoy bool
}

func (d Dog) Greet() string {
	return "Woof!"
}

func main() {
	d := Dog{
		Name:    "Wolfy",
		Breed:   "Husky",
		GoodBoy: true,
	}

	fmt.Println(d.Greet())
}
