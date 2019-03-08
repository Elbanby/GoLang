package main

import (
	"fmt"
	"reflect"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	path, err := config.Get("path")

	if err != nil {
		fmt.Println("Can't find key?! ", err)
	}

	enabled, err := config.GetBool("enabled")
	fmt.Println(reflect.TypeOf(path))
	fmt.Println(reflect.TypeOf(enabled))

	_, err = config.Get("nokey")
	if err != nil {
		fmt.Println(err)
	}

}
