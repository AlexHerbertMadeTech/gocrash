package main

import (
	"fmt"
)

func main() {
	ids := []int{4, 8, 15, 16, 23, 42}

	//loop

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID - %d\n", id)
	}

	// adding ids
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("total of all IDs = ", sum)

	//Range with map
	emails := map[string]string{"Lexie": "Lexie@Gmail.com", "Nadia": "Nadia@hotmail.com", "Ashleigh": "Ashleigh@whatthetrans.com"}
	for k, v := range emails {
		fmt.Printf("%s's email is %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s, email %s", k, emails[k])
	}

}
