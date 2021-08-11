package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("Don't need brackets; %d is less than %d\n", x, y)
	} else {
		fmt.Printf("can have brackets; %d is less than %d\n", y, x)
	}

	colour := "green"

	if colour == "red" {
		fmt.Println("colour is red")
	} else if colour == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is neither red nor blue")
	}

	switch colour {
	case "red":
		fmt.Println("I once knew a man")
	case "blue":
		fmt.Println("no I made that part up")
	default:
		fmt.Println("Hair! HAIR! HAAAAAIIIRRR!!!!!")
	}

}
