package main

import (
	"fmt"
	"strconv"
)

//definte person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method (value revceiver)
func (p Person) greet() string {
	return "Hello, I am " + p.firstName + " " + p.lastName + "! I am " + strconv.Itoa(p.age) + " years old"
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	//init person using struct
	Lexie := Person{
		firstName: "Alexandra",
		lastName:  "Herbert",
		city:      "Manchester",
		gender:    "Complicated",
		age:       35,
	}

	fmt.Println(Lexie.firstName)
	Lexie.hasBirthday()
	fmt.Println(Lexie.greet())
}
