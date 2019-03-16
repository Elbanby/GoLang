package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://github.com/elbanby")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", b)
}
