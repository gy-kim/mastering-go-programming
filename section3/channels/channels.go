package main

import "fmt"

func main() {
	// fmt.Println("//////// Channel //////////")
	// basicChannel()
	fmt.Println("//////// Buffered channel /////////")
	bufferedChannel()
	fmt.Println("//////// Closing channel /////////")
	closeChannel()
}

func closeChannel() {
	c := make(chan string)
	go SayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c
	fmt.Println("Channel close?", !ok, "value", v)
}

func SayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}

func bufferedChannel() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func basicChannel() {
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("Hello")

	// we send a signal to c in order to allow waitAndSay to continue
	c <- true

	// we wait to receive another sognal on c before we exit
	<-c
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}
