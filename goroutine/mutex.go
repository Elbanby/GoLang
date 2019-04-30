package main

import (
  "fmt"
  "sync"
)

func main() {
  var count int
  var lock sync.Mutex

  inc := func() {
    lock.Lock()
    defer lock.Unlock()
    count++
    fmt.Printf("Incrementing: %d\n",count)
  }

  dec := func() {
    lock.Lock()
    defer lock.Unlock()
    count--
    fmt.Printf("Decrementing: %d\n",count)
  }

  const numItterations = 5

  //Increment
  var arithmatic sync.WaitGroup
  for i := 0 ; i < numItterations ; i++ {
    arithmatic.Add(1)
    go func() {
      defer arithmatic.Done()
      inc()
    }()
  }

  // Decrement
  for i := 0 ; i < numItterations ; i++ {
    arithmatic.Add(1)
    go func() {
      defer arithmatic.Done()
      dec()
    }()
  }

  arithmatic.Wait()
  fmt.Println("All go routines are done!")
}
