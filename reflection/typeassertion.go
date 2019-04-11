package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		fmt.Printf("%T is a Stringer\n", b)
	}
	if i := 123; isStringer(i) {
		fmt.Printf("%T is a Stringer\n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
