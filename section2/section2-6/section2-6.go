package main

import "fmt"

type Set map[string]struct{}

func main() {
	s := make(Set)
	// add items
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	// get and print items
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
