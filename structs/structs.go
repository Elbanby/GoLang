package main

import (
	"fmt"
)

type subscriber struct {
	name      string
	age       int
	isPremium bool
}

func main() {
	user1 := newSubscriber("sam",25) // Returns a POINTER
	fmt.Printf("%#v\n", user1)
  user2 := newSubscriber("john",19)
  fmt.Printf("%#v\n", user2)
  // Make sure to pass the pointer to the struct
  upgradeToPremium(user2)
  fmt.Printf("%#v\n", user2)

  user3 := subscriber{"literalName", 20, false}
  fmt.Printf("%#v\n", user3)
}


func newSubscriber(name string, age int) *subscriber {
  var user subscriber
  user.name = name
  user.age = age
  user.isPremium = false
  return &user
}

func upgradeToPremium(user *subscriber) {
  user.isPremium = true
}
