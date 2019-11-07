package main

import "fmt"

var CAPACITY = 3

func main() {
	// Capacity of CAPACITY strings. Meaning the channel won't block until it receives 10 messages
	c := make(chan string, CAPACITY)

	c <- "hi"
	c <- "bye"
	close(c)

	for msg := range c {
		fmt.Println(msg)
	}
}