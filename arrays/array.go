package main

import (
  "fmt"
)

func main() {
  var notes [7]string
  notes[0] = "do"
  notes[1] = "re"
  fmt.Println(notes[0])
  fmt.Println(notes[1])
  fmt.Println(notes)

  var literals [3]string = [3]string{"el1","el2","el3"}

  for index, value := range literals {
    fmt.Println(index, value)
  }
}
