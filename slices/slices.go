package main

import "fmt"

func main() {
  var slice []string
  slice = make([]string, 10)
  slice[0] = "first item"
  fmt.Printf("this is a basic slice creation %s\n",slice[0])
  shortHandSlice := make([]string,11)
  fmt.Printf("this is a short hand notation for slice %s\n",shortHandSlice[0])
  fmt.Printf("first slice length is: %d and second slice length is: %d\n",len(slice),len(shortHandSlice))
}
