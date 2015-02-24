package main

import "fmt"

type age int

type user struct {
	Name  string
	Age   age
	Email string
}

// START AGE OMIT
func (a age) Generation() string {
	switch {
	case a < 35:
		return "millennial"
	case a < 55:
		return "generation X"
	case a < 75:
		return "baby boomer"
	default:
		return "greatest"
	}
}

// END AGE OMIT

func main() {
	// START MAIN OMIT
	u := user{
		Name:  "Ken Thompson",
		Age:   71,
		Email: "ken@example.com",
	}

	fmt.Println(u.Age.Generation())
	// END MAIN OMIT
}
