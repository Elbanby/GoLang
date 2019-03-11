package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(w, r)
			return
		} else if err != nil {
			fmt.Fprint(w, err)
		}
	}
	http.NotFound(w, r)
}

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	pr.Add("* /", home)
	http.ListenAndServe(":"+os.Getenv("PORT"), pr)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	fmt.Fprintln(w, "hello my name is: ", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
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
