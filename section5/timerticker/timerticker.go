package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("//////// Timer ///////")
	timer()

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
