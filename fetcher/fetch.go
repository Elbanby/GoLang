// Fetch prints the content found in a url
package main

import (
  "fmt"
  "os"
  "net/http"
  "io"
  "strings"
)


func main() {
  for _, url := range os.Args[1:] {

    if !strings.HasPrefix(url,"https://") && !strings.HasPrefix(url,"http://") {
       url = "https://" + url
    }

    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr,"fetch: %v\n", err)
      os.Exit(1)
    }
    b, err := io.Copy(os.Stdout,resp.Body)
    resp.Body.Close()

    if err != nil {
      fmt.Fprintf(os.Stderr,"fetch: %v\n", err)
      os.Exit(1)
    }
    fmt.Printf("%v\nstatus Code >>> %s\n resp heraders: >> %s", b, resp.Status, resp.Header)
  }
}
