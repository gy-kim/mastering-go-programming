package main

import "fmt"

func main() {
	basicSlice()
	refSlice()
}

func refSlice() {
	fmt.Println("")
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1)
	s2 := s1[2:4]
	s2[0] = 10
	fmt.Println("S1:", s1)
	fmt.Println("s2:", s2)
	s1[3] = 20
	fmt.Println("S2:", s2)
}

func basicSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]
	fmt.Println("Slices a ", a, " b:", b, " c:", c, " d:", d)
	fmt.Println("Capacity of b: ", cap(b))
	fmt.Println("What b actually sees: ", b[:len(b)])
}
