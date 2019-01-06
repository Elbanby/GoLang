//This is a cli guessing game
package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strings"
  "strconv"
  "time"
)

func main() {
  seconds := time.Now().Unix()
  rand.Seed(seconds)
  randnum := rand.Intn(100) + 1
  hasWon := false

  for i := 1 ; i <= 10 ; i++ {
    fmt.Printf("Guess a number: ")
    reader := bufio.NewReader(os.Stdin)
    guess, _ := reader.ReadString('\n')
    guess = strings.TrimSpace(guess)
    guessedValue, _ := strconv.Atoi(guess)


    if guessedValue  > randnum + 5 {
      fmt.Println("Oops! too high")
    } else if guessedValue < randnum - 5 {
      fmt.Println("Oops! too low")
    } else if guessedValue == randnum {
      fmt.Printf("\nYou win!! You got it in the %d guesses\n", i)
      hasWon = true
      break
    } else {
      fmt.Println("In range, continue guessing")
    }

    fmt.Printf("You have got %d gueese left\n\n", 10 - i)
  }

  if !hasWon {
    fmt.Printf("\nYou lose! Target was [%d]\n", randnum)
  }

}
