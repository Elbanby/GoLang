package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	// make sure not to specify deafault value for a bool opts
	Enabled bool   `short:"e" long:"enabled" description:"allows you to enable smth"`
	Name    string `short:"n" long:"name" default:"fakeName" description:"allow you to chnage the name"`
}

func main() {
	flags.Parse(&opts)
	if opts.Enabled {
		fmt.Println("Enabled")
	}
}
