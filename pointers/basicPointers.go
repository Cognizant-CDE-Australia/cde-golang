package main

import "fmt"

func main() {

	// Pointers  - No pointer arithmetic is invalid with golang
	var lastName *string = new(string)
	*lastName = "Smith"
	fmt.Println(*lastName)

	middleName := "Doe"
	fmt.Println(middleName)
	ptr := &middleName
	fmt.Println("Address = ", ptr, ", pointer value =", *ptr)

	// Attempting to change the value of the pointer variable
	middleName = "David"
	fmt.Println("Address = ", ptr, ", pointer value =", *ptr)

}
