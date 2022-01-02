package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = struct {
		sync.Mutex
		n int
	}{}
	workers := 1000

	ch := make(chan int, workers)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(ch chan<- int) {
			counter.Lock()
			defer counter.Unlock()
			counter.n += 1
			wg.Done()
		}(ch)
	}

	wg.Wait()
	fmt.Printf("The counter is: %v\n", counter.n)
}
