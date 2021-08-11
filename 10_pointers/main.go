package main

import "fmt"

func main() {
	a := 4825162342
	b := &a
	fmt.Println(a, b)
	fmt.Printf("Type is %T\n", b)

	//Change val with pointer
	*b = 10
	fmt.Println(*b)
}
