package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) GetCount() int {
	return cr.count
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) Intn(n int) int {
	fmt.Println("Outer Intn called...")
	cr.count++
	return cr.Rand.Intn(n) + 1
}

func main() {
	cr := NewCustomRand(time.Now().UnixNano())
	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())
}
