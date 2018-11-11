package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]
	fmt.Println("Slices a ", a, " b:", b, " c:", c, " d:", d)

	fmt.Println("Capacity of b: ", cap(b))
	fmt.Println("What b actually sees: ", b[:len(b)])
}
