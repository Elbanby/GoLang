package main

import (
	"fmt"
	"io"
	"reflect"
)

type Name struct {
	First, Last string
}

func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {
	n := &Name{"Inigo", "Montoya"}

	stringer := (*fmt.Stringer)(nil)
	implements(n, stringer)

	writer := (*io.Writer)(nil)
	implements(n, writer)
}

func implements(concrete interface{}, target interface{}) bool {
	fmt.Printf("The type of target is %T\n", target)
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	fmt.Printf("The value of Concrete: %v\n", v)

	t := reflect.TypeOf(concrete)
	fmt.Printf("Typeof Concrete is: %v\n", t)

	if t.Implements(iface) {
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("%T is not %s\n", concrete, iface.Name())
	return false
}
