package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "food"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println("strings.EqualFold", strings.EqualFold("Go", "go"))

	s := []string{"food", "bar", "baz"}
	fmt.Println("strings.Join", strings.Join(s, ", "))
}
