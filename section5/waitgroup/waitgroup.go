package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		// Increment the WaitGroup counter
		wg.Add(2)
		// Launch a goroutine
		go func(i int) {
			// Decrement the couter when the goroutine completes.
			defer wg.Done()
			// Do some work
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for ", i)
		}(i)
	}

	wg.Wait()
}
