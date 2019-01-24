package main

import (
  "fmt"
  "log"
)

type OverHeatError float64

// This satisfies the error interface
func (o OverHeatError) Error() string {
    return fmt.Sprintf("Over heating by %.2f degrees!", o)
}

func CheckTemprature(actual float64, safe float64) error {
  excess := actual - safe
  if excess > 0 {
    return OverHeatError(excess)
  }
  return nil
}

func main() {
   err := CheckTemprature(10,5)
  if err != nil {
    log.Fatal(err)
  }
}
