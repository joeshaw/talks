package main

import "fmt"

// START AGE METHOD OMIT
// START AGE OMIT
type age int

// END AGE OMIT
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

// END AGE METHOD OMIT

// START MAIN OMIT
type user struct {
	Name  string
	Age   age
	Email string
}

func main() {
	u := user{
		Name:  "Ken Thompson",
		Age:   71,
		Email: "ken@example.com",
	}

	fmt.Println(u.Age.Generation())
}

// END MAIN OMIT
