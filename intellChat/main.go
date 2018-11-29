package main

import (
  "log"
  "net/http"
)

// templ represents a single template
type templateHandler struct {
  once sync.Once
  filename string
  temp1 *template.Template
}

// ServeHTTP handles the HTTP request.

func (t *templateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  t.once.Do(func () {
    t.temp1  = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, nil)
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`
      `))
  })

   // start the web server
   if err := http.ListenAndServe(":8080",nil); err != nil {
     log.Fatal("ListenAndServe:", err);
   }
}
