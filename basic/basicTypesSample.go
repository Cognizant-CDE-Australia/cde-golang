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
}
