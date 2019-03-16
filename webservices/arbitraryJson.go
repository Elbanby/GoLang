package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var jsonData = []byte(`{
"firstName": "Jean",
"lastName": "Bartik",
"age": 86,
"education": [
 {
 "institution": "Northwest Missouri State Teachers College",
 "degree": "Bachelor of Science in Mathematics"
 },
 {
 "institution": "University of Pennsylvania",
 "degree": "Masters in English"
 }
],
"spouse": "William Bartik",
"children": [
 "Timothy John Bartik",
 "Jane Helen Bartik",
 "Mary Ruth Bartik"
]
}`)

func main() {
	var j = make(map[string]interface{})
	err := json.Unmarshal(jsonData, &j)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// strings
	fmt.Println(j["firstName"])
	fmt.Println(j["spouse"])
	//
	fmt.Println(j["age"])

	// Array of objects
	for _, v := range j["education"].([]interface{}) {
		if value, ok := v.(map[string]interface{}); ok {
			fmt.Printf("\nInstution: %s\nDegree: %s\n\n", value["institution"], value["degree"])
		}
	}

	// array of strings
	for _, v := range j["children"].([]interface{}) {
		fmt.Printf("children : %s\n", v)
	}

	fmt.Println("\n\n\nOr use the printJson function on this file")
	printJSON(j)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}
