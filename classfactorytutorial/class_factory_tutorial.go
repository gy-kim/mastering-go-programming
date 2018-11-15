package main

import (
	"fmt"

	"github.com/gy-kim/mastering-go-programming/classfactorytutorial/appliances"
)

func main() {
	// Use the class factory to create an appliance of the requested type
	// Request the user to enter the appliance type
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	// use fmt.scan to retrive the user's input
	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	// if no errors start the appliance then print it's purpose
	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		// if error encounted, print the error
		fmt.Println(err)
	}
}
