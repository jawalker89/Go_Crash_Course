package main

import "fmt"

func main() {
	//var name string = "Jason"
	//Shorthand way to declare a variable. Can only be used inside a function.
	name := "Jason"
	size := 1.5
	var age int32 = 47
	const isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)

}
