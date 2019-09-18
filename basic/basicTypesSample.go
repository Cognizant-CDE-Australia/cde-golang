package main

import (
	"fmt"
)

func main() {
	
	var i int
	i = 10
	fmt.Println(i)

	var f float32 = 91.8
	fmt.Println(f)

	//Implicit initialization
	firstName := "John"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(4, 8)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// Pointers  - No pointer arithmetic is invalid with golang
	var lastName *string = new(string)
	*lastName = "Smith"
	fmt.Println(*lastName)

	middleName := "Doe"
	fmt.Println(middleName)
	ptr := &middleName
	fmt.Println(ptr, *ptr)

	// Attempting to change the value of the pointer variable
	middleName = "David"
	fmt.Println(ptr, * ptr)
	
}
