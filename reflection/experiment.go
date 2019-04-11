package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Test struct {
	name string
	age  int
}

// func (t *Test) FakeMethod() {
//
// }

func main() {
	p := &Test{"omar", 10}

	// Pointer
	t := reflect.ValueOf(p)
	fmt.Println(t.Kind())
	fmt.Println(t.Type())

	// reflect Indirect derefrences the Pointer
	t1 := reflect.Indirect(reflect.ValueOf(p))
	fmt.Println(t1)
	fmt.Println(t1.Kind())
	fmt.Println(t1.Type())

	// strings Repeat repeats whatever string you provide
	tripleS := strings.Repeat("s", 3)
	fmt.Println(tripleS)

	fmt.Println("\nSome handy methods with structs types")
	fmt.Printf("Number of fields in %q %s are %d\n", t1.Type(), t1.Kind(), t1.NumField())

	tType := t1.Type()
	tValue := reflect.Indirect(t1.Field(0))
	fmt.Println(tType.Field(0))
	fmt.Println(tValue)

	// fmt.Printf("Accessing the first field %v\n", t1.Field(1))
}
