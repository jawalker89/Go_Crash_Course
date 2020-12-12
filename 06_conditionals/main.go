package main

import "fmt"

func main() {
	x := 10
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	colour := "green"

	if colour == "red" {
		fmt.Println("Colour is red")
	} else if colour == "blue" {
		fmt.Println("Colour is blue")
	} else {
		fmt.Println("Colour is NOT red or blue")
	}

	// Switch
	switch colour {
	case "red":
		fmt.Println("Colour is red")
	case "blue":
		fmt.Println("Colour is blue")
	default:
		fmt.Println("Colour is NOT red or blue")
	}
}
