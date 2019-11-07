package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sender(1000, c1, "Message from channel 1")
	go sender(500, c2, "Message from channel 2")

	for {
		select {
		case msg := <- c1:
			fmt.Println(msg)
		case msg := <- c2:
			fmt.Println(msg)
		}
	}

}

func sender(duration int, c chan string, msg string) {
	for {
		c <- msg
		time.Sleep(time.Millisecond * time.Duration(duration))
	}
}