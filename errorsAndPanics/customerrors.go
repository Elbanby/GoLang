package main

import (
	"fmt"
)

func main() {
	if err := lint(); err != nil {
		fmt.Println(err.Error())
	}
}

type Myerror struct {
	Message     string
	Linenumber  int
	Suggesstion string
}

func (m *Myerror) Error() string {
	return fmt.Sprintf("Error: %s, on line number: %d", m.Message, m.Linenumber)
}

func lint() error {
	err := Myerror{"my custom error message", 2, "I suggest you handle this err!"}
	return &err
}
