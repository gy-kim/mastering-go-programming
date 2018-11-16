package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("//////// Timer ///////")
	// timer()
	fmt.Println("//////// Cancel timer outside go routine ///////////")
	cancelTimer()
}

func cancelTimer() {
	nc := make(chan int)
	stopc := make(chan bool)

	go CancelSlowCounter(1, nc, stopc)
	time.Sleep(5 * time.Second)

	nc <- 2
	time.Sleep(6 * time.Second)
	stopc <- true
	time.Sleep(1 * time.Second)
}

func CancelSlowCounter(n int, nc chan int, stopc chan bool) {
	i := 0
	// Create a duration of n seconds
	d := time.Duration(n) * time.Second

Loop:
	for {
		select {
		// Use time.After channel to wait for a time period
		case <-time.After(d):
			i++
			fmt.Println(i)
		case n = <-nc:
			fmt.Println("Timer duration changed to", n)
		case <-stopc:
			fmt.Println("Timer stopped")
			break Loop
		}
	}

	fmt.Println("Exiting slow Counter")
}

func timer() {
	go SlowCounter(2)
	time.Sleep(15 * time.Second)
}

func SlowCounter(n int) {
	i := 0
	// create a duration of n seconds
	d := time.Duration(n) * time.Second

	for {
		// Create a timer for this duration

		// t := time.NewTimer(d)
		// <-t.C
		<-time.After(d)
		i++
		fmt.Println(i)
	}
}
