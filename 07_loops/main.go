package main

import "fmt"

func main() {
	// long method

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	//Fizz buzz
	for i := 1; i <= 100; i++ {
		// fmt.Printf("Number: %d\n", i)
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if i%5 != 0 && i%3 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
