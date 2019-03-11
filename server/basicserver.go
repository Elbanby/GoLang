package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	// this path wil cause redirection issues and will drop query strings
	http.HandleFunc("/goodbye/", bye)
	http.HandleFunc("/", home)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	fmt.Fprintln(w, "hello my name is: ", name)
}

func bye(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "default name"
	}
	fmt.Fprintln(w, "Goodbye: ", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Home page welcomes you")
}
