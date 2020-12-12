package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// has Birthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {
	//Initialise person using struct
	person1 := Person{firstName: "Samamtha", lastName: "Smith", city: "Boston", gender: "Female", age: 25}

	// Alternative syntax
	person2 := Person{"Bob", "Johnson", "New York", "Male", 30}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")

	fmt.Println(person2.greet())

}
