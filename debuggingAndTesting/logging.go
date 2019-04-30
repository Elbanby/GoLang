package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	// io.Writer, prefix, flags
	logger := log.New(logFile, "logger: ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular log message")
	logger.Fatalln("This is a fatal error")
}
