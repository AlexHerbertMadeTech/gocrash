package main

import "fmt"

func main() {
	//Arrays
	var FruitArr [2]string

	//Assign
	FruitArr[0] = "mango"
	FruitArr[1] = "banana"

	//Declare and assign
	vegArr := [2]string{"Carrot", "Leek"}

	fmt.Println(FruitArr)
	fmt.Println(vegArr)

	breadSlice := []string{"White", "Brown", "Seeded"}

	fmt.Println(breadSlice)
}
