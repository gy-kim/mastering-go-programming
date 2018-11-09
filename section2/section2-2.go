package main

import (
	"fmt"
	"runtime"
)

func main() {
	var v1 int // declare a variable
	v1 = 1
	var v2 = 2
	v3 := 3
	fmt.Println(v1 + v2 + v3)

	// A pointer holds the menory address of a variable

	var p1 *int // declare a pointer
	i := 4      // create an integer using type inference
	p1 = &i     // assign the address of the integrer to the pointer
	fmt.Println(p1)

	p2 := &i         // create a pointer via type inference
	fmt.Println(*p2) // use the * in order to 'dereference' the pointer and get the original value

	n := 32
	f := float64(n)
	const pi = 3.14
	fmt.Println(pi * f)

	// arrays
	var a [2]int
	a[0], a[1] = 4, 5
	fmt.Println("Array a", a, len(a))

	b := [3]int{9, 8, 10}
	fmt.Println("Array a", a, len(b))

	// a slice
	c := []int{99, 44}
	c = append(c, 84)
	fmt.Println("Slice c", c, len(c))

	// functions pass by value, so we need the argument to be a pointer if we want to change the value of the passed variable
	X := 5
	ChangeX(&X)
	fmt.Println("Value of X:", X)

	// function value
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(computeMultiplyVal(2, add))

	// function closer
	inc := incrementer()
	fmt.Println(inc(), inc(), inc())

	// if statement
	if i := inc(); i < 0 {
		fmt.Println("i is negative number", i)
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is a positive number", i)
	}

	// switch statements
	switch i := inc(); {
	case i < 0:
		fmt.Println("i is a negative number", i)
	case i == 0:
		fmt.Println("i is zero")
	default:
		fmt.Println("i is a positive number", i)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows..
		fmt.Printf("%s.", os)

		defer fmt.Println("Exiting functions.... ")
		fmt.Println("Entering function")
	}

}

const (
	big   = 1 << 100
	small = big >> 99
)

func computeMultiplyVal(val int, fn func(a, b int) int) int {
	return val * fn(val, val)
}

func ChangeX(X *int) {
	*X = 10
}

func incrementer() func() int {
	// initializes i
	i := 0
	return func() int {
		// then new value of it will be retained
		i++
		return i
	}
}
