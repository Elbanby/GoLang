/*
	This is a very dense example. Its purpose is to
	walk through any struct and prints all its members and their
	types.

	Check "experiment.go" to clarify some of the concepts
	here!
*/
package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MyInt int

type Person struct {
	Name      *Name
	Address   *Address
	NickNames []string
}

type Name struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}

func main() {
	fmt.Println("Walking a simple integer")
	var one MyInt = 1
	walk(one, 0)
	fmt.Println("")

	two := struct{ Name string }{"foo"}
	walk(two, 0)
	fmt.Println("")

	p := &Person{
		Name:      &Name{"mr", "john", "green"},
		Address:   &Address{"Fake street", "bay area"},
		NickNames: []string{"johnny"},
	}
	walk(p, 0)

}

func walk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()
	tabs := strings.Repeat("\t", depth+1)
	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := reflect.Indirect(val.Field(i))
			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is type %q (%s)\n", tabs, field.Name, field.Type, fieldValue.Kind())
			if fieldValue.Kind() == reflect.Struct {
				walk(fieldValue.Interface(), depth+1)
			}
		}
	}

}
