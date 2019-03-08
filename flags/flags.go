package main

import (
	"flag"
	"fmt"
)

// A pointer to a variable. This will createa a flag -name
var username = flag.String("name", "default-name", "A default flag description")
var phone int
var weight float64
var help bool
var pretty bool

// This function should be there. I believe it is called by the flag pacakge
func init() {
	flag.BoolVar(&help, "help", false, "display the help page")
	flag.BoolVar(&help, "h", false, "display the help page")

	flag.BoolVar(&pretty, "pretty", false, "Show help page in a prettier format")
	flag.BoolVar(&pretty, "P", false, "Show help page in a prettier format")

	flag.IntVar(&phone, "phone", 000000000, "Allows user to enter phone as an int")
	flag.Float64Var(&weight, "weight", 0.0, "double value for weight")
}

func main() {
	// if flag parge is not called the flag will have its default variable
	flag.Parse()
	fmt.Println(*username)

	if help == true && pretty == false {
		flag.PrintDefaults()
	}

	if pretty == true {
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("-%v: %v\n", f.Name, f.Usage)
		})
	}

	//When you initalize the flags with the init function you can't access the vars as pointers
	// meaning *help will fail!
	fmt.Println(help)
	fmt.Println(phone)
	fmt.Println(weight)
}
