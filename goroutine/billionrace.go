/*
 * WaitGroup will help on waiting till all go routines are done
 * Mutex will lock the memory to one go routine at a time.
 * This programs prevents a billion race condition!
 */
package main

import (
	"fmt"
	"sync"
)

var secretnumber int64 = 500

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000000000; i++ {
		wg.Add(2)
		go increment(&wg, &m)
		go decrement(&wg, &m, 2)
	}

	wg.Wait()
	fmt.Printf("Result is %v\n", secretnumber)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	secretnumber++
	m.Unlock()
}

func decrement(wg *sync.WaitGroup, m *sync.Mutex, decreaseBy int64) {
	defer wg.Done()
	m.Lock()
	secretnumber -= decreaseBy
	m.Unlock()
}
