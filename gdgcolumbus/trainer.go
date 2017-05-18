package main

import "fmt"

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
	return fmt.Sprintf("Hello %s", p.Name)
}

// START OMIT

type Trainer struct {
	Person
	Experience int
}

func main() {
	t := Trainer{
		Person:     Person{Name: "Helga"},
		Experience: 20,
	}

	fmt.Println(t.Greet())
}

// END OMIT
