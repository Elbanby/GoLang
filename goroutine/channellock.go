package main

import (
	"fmt"
	"time"
)

func main() {
	// make a buffer channel with one space
	lock := make(chan bool, 1)
	// start go routines, sharing the same lock
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

/*
  A worker acquires the lock by sending
  it a message. The first worker to hit
  this will get the one space, and thus
  own the lock. The rest will block.
*/
func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	// add a message to the channel hence block
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	// Relase the lock by reading(consuming) its data
	<-lock
}
