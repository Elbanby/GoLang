package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	// starts a server listing on port 3000
	listner, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Failed to open on port 3000")
		return
	}
	// Listen for new connections and handle any connection errors
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			return
		}
		// when a connection is accepted pass it to a handle routine
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// creating a panic handler. Note that go routines when panic
	// they keep unwinding, but not to the src that created the routine
	// rather to the begining of the go routine stack.
	// stops the panic from propagating up the stack
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		conn.Close()
	}()
	// tries to reaed a line of data from connection
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	// we panic to demonstare goroutine unwinding
	panic(errors.New("Error occured after connection closed"))
}
