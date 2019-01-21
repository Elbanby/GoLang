package main

import (
  "fmt"
  "sort"
)


func main() {
  //Arrays
  var array [10]string = [10]string{"hi","there","sugar","bear"}
  fmt.Println(array[0])

  //Slices
  var slice []string
  slice = make([]string,10)
  fmt.Println(slice[0])

  // SHort variable declartion
  sliceShort := make([]string,10)
  fmt.Println(sliceShort[0])

  // Creating an empty slice
  emptySlice := []string{}
  emptySlice = append(emptySlice,"HI from slice")
  fmt.Println(emptySlice[0])

  // Declare a map
  var myMap map[string]float64
  myMap = make(map[string]float64, 10)
  myMap["i0"] = 10.32423402
  fmt.Println(myMap["i0"])

  // Short variable decleration
  myMapShort := make(map[string]string,10)
  myMapShort["li"] = "lithium"
  fmt.Println(myMapShort["lo"])
  fmt.Println(myMapShort["li"])

  // Creating an empty map without using make
  myEmptyMap := map[string]int{}
  myEmptyMap["fakePerson"] = 416415416
  fmt.Println(myEmptyMap["fakePerson"])

  // chechking if a map value is the zero value or true value
  grades := map[string]float64{"Jhon": 75, "Peter": 65.5, "Jordan": 80, "Harmet": 95, "Jin": 98, "Andrew": 96.5}
  value, hasAssigned := grades["Jhon"]
  fmt.Printf("John grade is:%.1f and has hasAssigned is: %v\n",value, hasAssigned)
  value, hasAssigned = grades["Sam"]
  fmt.Printf("Sam grade is:%.1f and has hasAssigned is: %v\n",value, hasAssigned)

  // To delete a key from a map
  value, hasAssigned = grades["Peter"]
  fmt.Printf("Peter grade is:%.1f and has hasAssigned is: %v\n",value, hasAssigned)

  delete(grades,"Peter")

  value, hasAssigned = grades["Peter"]
  fmt.Printf("Peter grade is:%.1f and has hasAssigned is: %v\n",value, hasAssigned)

  fmt.Println(len(grades))

  //Note that maps access keys randomly this line will print different results each time
  for key, value := range grades {
    fmt.Printf("%0s grade is: %4.1f\n",key, value)
  }
  fmt.Println("-----------------------")
  // to handle that situation we will have to do it programatically
  keysSlice := make([]string, len(grades))
  i := 0
  for key := range grades {
    keysSlice[i] = key
    i++
  }
  // now lets sort keysSlice Alphabitically
  sort.Strings(keysSlice)
  // Now lets loop through the map again, but this time we specify thr sorted keys
  for _, name := range keysSlice {
    fmt.Printf("%s grades is \t%2.1f\n",name, grades[name])
  }

}
