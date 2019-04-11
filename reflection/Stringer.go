package main

import (
	"fmt"
)

type UserInfo struct {
	Name string
	Age  int
}

func (u *UserInfo) String() string {
	return fmt.Sprintf("Your Struct of type %T has a name prop of value %s", u, u.Name)
}

func main() {
	user := &UserInfo{"john", 20}
	fmt.Println(user)
}
