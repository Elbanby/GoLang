package main

import (
	"fmt"
	"sync"
)

var Stringer = make(chan string, 2)
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		ChanWritter(Stringer, "first msg")
	}()

	go func() {
		defer wg.Done()
		ChanReader(Stringer)
	}()

	wg.Wait()
	fmt.Println("Done...")
}

// A write only channel
func ChanWritter(ch chan<- string, msg string) {
	ch <- msg
}

// A read only channel
func ChanReader(ch <-chan string) {
	fmt.Println("Channel recieved a new message: " + <-ch)
}
