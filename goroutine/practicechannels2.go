package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 2)
var doneSignal = make(chan struct{})

func main() {
	go channelListener()
	strChan <- "Msg1"
	strChan <- "Msg2"
	time.Sleep(100 * time.Millisecond)
	doneSignal <- struct{}{}
	fmt.Println("Done! channel closed!")
	close(strChan)
}

func channelListener() {
	for {
		select {
		case str := <-strChan:
			fmt.Println("Channel recieved a new message " + str)
		case <-doneSignal:
			break
		}
	}
}
