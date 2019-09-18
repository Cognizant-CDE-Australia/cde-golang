package main

import "fmt"

//Anonymous function
var add = func(a, b int) int {
	return a+b
}

func main() {
	fmt.Println("5+3 = ", add(5,3))
}
