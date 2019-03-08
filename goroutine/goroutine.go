package main

import (
  // "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {
  defer wg.Wait()
  wg.Add(5)
  go multiplication(2)
  go multiplication(3)
  go multiplication(4)
  go multiplication(5)
  go multiplication(6)

  // multiplication(2)
  // multiplication(3)
  // multiplication(4)
  // multiplication(5)
  // multiplication(6)
}

func multiplication(table int) {
  defer wg.Done()
  for i := 0 ; i < 12 ; i++ {
    // fmt.Printf("%v x %v = %v\n",i, table, i *table)
  }
}
