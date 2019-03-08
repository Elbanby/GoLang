package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.facebook.com")

	if err != nil {
		log.Fatal("An error has occured.. exit")
	}

	if resp.StatusCode == http.StatusNotFound {
		log.Fatal("404")
	}

	respData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData), "\n\n\n\n")

	fmt.Println(resp.Header)
}
