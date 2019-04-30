package main

import (
  "fmt"
  "sync"
)

func main() {
/* example one begins
  wg.Add(1)
  go Test()
  wg.Add(1)
  go func() {
    defer wg.Done()
    fmt.Println("2nd go routine")
  }()
  wg.Wait()
  fmt.Println("All routines are done")
  */

  hello := func(wg *sync.WaitGroup, id int) {
    defer wg.Done()
    fmt.Printf("Hello from %v!\n",id)
  }

  const numGreetings = 5
  var wg sync.WaitGroup
  wg.Add(numGreetings)
  for i := 0 ; i < numGreetings ; i++ {
    go hello(&wg,i + 1)
  }
  wg.Wait()
  fmt.Println("Done with all go routines!")
}

// func Test() {
//   defer wg.Done()
//   fmt.Println("1st go routine! hi from test function!")
// }
