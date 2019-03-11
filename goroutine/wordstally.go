package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type words struct {
	sync.Mutex
	Found map[string]int
}

func newWords() *words {
	return &words{make(map[string]int)}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.Found[word]
	if !ok {
		w.Found[word] = n
		return
	}
	w.Found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	var wg sync.WaitGroup
	dict := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if err := tallyWords(file, dict); err != nil {
				fmt.Println(err.Error())
			}
		}(f)
		wg.Wait()
		fmt.Println("Words that appears more than once")
		for word, count := range dict.Found {
			if count > 1 {
				fmt.Printf("%s: %d\n", word, count)
			}
		}
	}

}
