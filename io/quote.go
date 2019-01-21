// This program prints rnadom quotes from a text file
package main

import (
  "github.com/elbanby/ioString"
  "math/rand"
  "time"
  "fmt"
  "log"
)

func main() {
  filedata, err := ioString.Read("quotes.txt")
  quotes := make([]string,0)
  errLog(err)

  for _, quote := range filedata {
    // add the text to the slice
    quotes = append(quotes,quote)
  }
  fmt.Println("Here is your quote of the day\n")
  index := randIndex(len(quotes))
  fmt.Println(quotes[index])
}

func errLog(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func randIndex(len int) int {
  // now create a random num between 0 and the slice length
  seconds := time.Now().Unix()
  rand.Seed(seconds)
  return rand.Intn(len)
}
