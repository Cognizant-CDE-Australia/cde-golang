package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before ", &x)
	fmt.Println("x before ", x)

	populateX(&x)

	fmt.Println("x after ", &x)
	fmt.Println("x after ", x)
}

func populateX(y *int) {
	fmt.Println("y before ", y)
	fmt.Println("y before ", *y)

	*y = 43

	fmt.Println("y after ", y)
	fmt.Println("y after ", *y)
}
