// passfail will report if a student will pass or fail
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
  "strconv"
)

func main() {
  fmt.Print("Enter a grade: ")
  reader := bufio.NewReader(os.Stdin)
  input , err := reader.ReadString('\n')
  if (err != nil) {
    log.Fatal(err)
  }
  input = strings.TrimSpace(input)
  grade, err := strconv.ParseFloat(input,64)
  if (err != nil) {
    log.Fatal(err)
  }
  var status string
  if grade >= 60 {
    status = "passing"
  } else {
    status = "failing"
  }
  fmt.Println(status)
}


/*
  if a function returns more than one value and we are assigning
  the value to one variable we get an error:
  multiple-value <function name> in single-value context

  This is importnat to be aware of
*/
