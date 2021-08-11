package main

import "fmt"

func main() {
	name := "Nadia"
	var age int = 35
	name = "Alex"
	fmt.Printf("Hello, %v, you are %v years old!\n", name, age)
	fmt.Printf("Age is a %T", age)
}
