package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func main() {
  fmt.Print("Enter file path: ")
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if (err != nil) {
    log.Fatal(err)
  }
  fileInfo, error := os.Stat(strings.TrimSpace(input))
  if error != nil {
    log.Fatal(error)
  }
  fmt.Printf("file %s size is %d bytes\n", strings.TrimSpace(input), fileInfo.Size())
}
