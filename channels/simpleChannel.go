package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string)
	go sender(10,c)

	// If you don't use the range this will terminate after one message
	for msg := range c {
		fmt.Println(msg)
	}
}

func sender(iterations int, c chan string) {
	for i := 0; i <= iterations; i++  {
		c <- "hello - " + strconv.Itoa(i)
	}
	// Must close the channel after finishing using it
	close(c)
}


