package main

import (
	"flag"
	"fmt"
	"sync"
)

var amount = flag.Int("amount", 10, "amount of goroutines")

func main() {
	flag.Parse()
	fmt.Printf("Program is working...\n__amount of set goroutines = %v\n", *amount)

	var finished = struct {
		sync.Mutex
		count int
	}{}
	wg := sync.WaitGroup{}
	wg.Add(*amount)

	for i := 0; i < *amount; i++ {
		go func() {
			finished.Lock()
			defer finished.Unlock()
			finished.count += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("__amount of finished goroutines = %v\nProgram has finished\n", finished.count)
}
