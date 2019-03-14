package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Println(result)
	}
}

// Concat taks an arbitrary number of arguments and concat them
// If concat gets called with no string it will return an empty space and an error
func Concat(strs ...string) (string, error) {
	if len(strs) == 0 {
		return "", errors.New("No string supplied")
	}
	return strings.Join(strs, " "), nil
}
