package main

import (
	"fmt"
	"reflect"
)

func main() {
	n1, n2 := 2, 10.1
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(42.00))
	fmt.Println(reflect.TypeOf(42.000000000000000000001))
	fmt.Println(reflect.TypeOf("st"))
	fmt.Println(reflect.TypeOf('O'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println("%d %d", n1*int(n2))
}
