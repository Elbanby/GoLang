package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// type switches operate on types (not kinds)
type myInt int64

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	var d myInt = 1
	res := sum(a, b, c, d)
	fmt.Printf("Result is %.3f\n", res)
}

func sum(v ...interface{}) float64 {
	var res float64
	for _, val := range v {
		ref := reflect.ValueOf(val)

		switch ref.Kind() {
		case reflect.Int:
			res += float64(ref.Int())
		case reflect.Uint8:
			res += float64(ref.Uint())
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsporrted type %T. Ignoring. \n", val)
		}
	}
	return res
}
