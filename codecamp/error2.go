package main

type errorString string

func (es errorString) Error() string {
	return es
}

func main() {
	str := "this talk makes no sense"

	var someError error
	someError = str
}
