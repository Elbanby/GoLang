package main

import (
  "fmt"
  "reflect"
)

type Liters float64
type Gallons float64
type Testme string

func (t Testme) test() {
  fmt.Println(t)
}

func main() {
  var usCar Gallons
  var otherCars Liters

  usCar = 10.0
  otherCars = 200
  fmt.Println(usCar)
  fmt.Println(otherCars)

  var num float64
  num = 12.09

  usCar = Gallons(num)
  fmt.Println(usCar)

  usCar = Gallons(otherCars) // legal but incorrect
  fmt.Println(usCar)

  var str Testme
  str = Testme("I am sent from test")
  fmt.Println(reflect.TypeOf(str))

}
