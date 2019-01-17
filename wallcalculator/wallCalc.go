//Calculate paint needed to paint a wall
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  numWalls := getvalue("walls")
  var buildStr string

  for i := 0 ; i < toNumber(numWalls); i++ {
      width := getvalue("width")
      height := getvalue("height")

      widthVal, _ := strconv.ParseFloat(width,64)
      heightVal, _ := strconv.ParseFloat(height,64)

      buildStr += fmt.Sprintf("wall %d | %2.2f liters\n", i,  paintNeeded(widthVal,heightVal))
  }

  fmt.Println(buildStr)
}

func paintNeeded(width float64, height float64) float64 {
  return width * height / 10
}


func getvalue(valName string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Eneter %s: ",valName)
    tempStr, _ := reader.ReadString('\n')
    tempStr = strings.TrimSpace(tempStr)
    fmt.Println("")
    return tempStr
}

func toNumber(num string) int {
  numVal, _ := strconv.Atoi(num)
  return numVal
}
