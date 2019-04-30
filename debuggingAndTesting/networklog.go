package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902 ")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "network", f)
	logger.Println("This is a regular log message")

	/*
		"notice that we changed log.Fatalln to a log.Panicln in this example?
		There’s a simple reason for this; the log.Fatal* functions have an unfortunate
		side effect: the deferred function isn’t called. Why not? Because log.Fatal* calls
		os.Exit, which immediately terminates the program without unwinding the function
		stack."
	*/
	logger.Panicln("Error log message - halting")
}
