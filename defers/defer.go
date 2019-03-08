/*
  Multiple defer statements are executed in the reverse order; last-in-first-out
  or LIFO.
  Arguments passed to deferred functions are evaluated immediately
*/

package main

import (
  "fmt"
)

func main() {
  defer fmt.Println("first")
  defer fmt.Println("second")
  defer fmt.Println("third")

  sum := new(int)
  fmt.Println(*sum)
}
