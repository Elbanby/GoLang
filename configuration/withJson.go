package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("conf.json")
	defer file.Close()

	if err != nil {
		fmt.Println("Erron: ", err)
	}

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}
