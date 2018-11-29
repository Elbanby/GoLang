// Dup1 prints the text of each line that appears more than once in the
// standard input and the number of occurrences
package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  //IDK as of now
  counts := make(map[string]int)
  //To read from stundered input
  input := bufio.NewScanner(os.Stdin)

  for input.Scan() {
    counts[input.Text()]++
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
