package main

import "fmt"

func main() {
	//define map
	emails := make(map[string]string)

	// Assign k+v
	emails["Bob"] = "bob@gmail.com"
	emails["Lexie"] = "Lexie@gmail.com"
	emails["Ashleigh"] = "Ashleigh@whatthetrans.com"

	fmt.Println(emails["Ashleigh"])
	fmt.Printf("There are %d emails", len(emails))

	delete(emails, "Bob")
	fmt.Println(emails)
}
