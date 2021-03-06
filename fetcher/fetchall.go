// Fetch urls concurrently. Hence time taken depends on the time taken for all requests to return
package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url,ch) //Start go routine!
  }
  for range os.Args[1:] {
    fmt.Println(<-ch) //Recieve from channel ch
  }
  fmt.Println("%.2fs elapased\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // Send to channel ch
    return
  }

  nbytes , err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close() // don't leak resources
  if err != nil {
    ch <- fmt.Sprint("while reading %s: %v", url, err)
    return
  }

  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
