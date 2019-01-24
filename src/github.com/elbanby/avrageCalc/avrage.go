package main

import "fmt"

func main() {
  var avrage float64
  nums := [10]float64{1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0}

  for _, value := range nums {
    avrage += value
  }

  avrage = avrage / float64(len(nums))
  fmt.Println(avrage)
}
