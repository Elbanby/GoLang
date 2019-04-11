package main

import (
	"fmt"
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
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case uint8:
			res += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		// case myInt:
		// 	res += float64(val.(myInt))
		default:
			fmt.Printf("Unsporrted type %T. Ignoring. \n", val)
		}
	}
	return res
}
