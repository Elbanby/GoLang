// Echo prints its command-line arguments.
package main

import (
        "fmt"
        "os"
        "time"
)

// func main() {
//   var s, sep string
//
//   for i := 1 ; i < len(os.Args) ; i++ {
//     s += sep + os.Args[i]
//     sep = " "
//   }
//   fmt.Println(s)
// }

// func main() {
//   s, sep := "", " "
//   for _, arg := range os.Args[1:] {
//     s += sep + arg
//     sep = " "
//   }
//   fmt.Println(s)
// }


// func main() {
//   fmt.Println(os.Args[1:])
// }


func main() {
  start := time.Now()
  fmt.Print("excercise questions:\n1-")
  s, sep := "" , " "
  for _, args := range os.Args[0:] {
    s += s + sep + args
  }
  fmt.Println(s)

  fmt.Println("excercise 2")

  for i, args := range os.Args[0:] {
    fmt.Println(i, args)
  }

  t := time.Now()

  fmt.Println("start: ", start, " \nEnd: ", t)
}
