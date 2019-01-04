package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(strings.Title("using some of the lessons learned in c1"))
	fmt.Println("first- Simple arithmatics and type conversions")
	var n1, n2 int = 20, 22
	fmt.Println("the sum is %d", n1+n2)
	x, y := 2, 2.0
	fmt.Println("the sum is %d", math.Floor(float64(x)+y))
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))
}
