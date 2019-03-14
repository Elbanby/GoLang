package main

import "errors"

func main() {
	// Panic examples
	// panic("you can pass a message")
	// panic(nil)
	panic(errors.New("something bad happened"))
}
