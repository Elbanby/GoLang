package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  str := ""
  for _, filename := range os.Args[1:] {
    str  += filename + " "
    data, err := ioutil.ReadFile(filename)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
      continue
    }
    for _, line := range strings.Split(string(data),"\n") {
      counts[line]++
    }
  }

  for line, n := range counts {
    fmt.Printf("%s \t %d \t %s\n", str,n, line)
  }
}
