package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
)

func main() {
	name, _ := os.Hostname()
	fmt.Println(name)

	addrs, _ := net.LookupHost(name)
	fmt.Println(addrs)
	fmt.Println(reflect.TypeOf(addrs))
}
