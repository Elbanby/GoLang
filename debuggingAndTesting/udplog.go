package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("udp", "localhost:1902", 30*time.Second)
	if err != nil {
		panic("Error occured establishing the connection")
	}
	defer conn.Close()

	logger := log.New(conn, "logger: ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular log message")
	logger.Panicln("Error log message - halting")
}
