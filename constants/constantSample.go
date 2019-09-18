package main

import "fmt"

const (
	first = iota + 3
	second
)

func main() {
		const value = 45 //Implicitly typed constant - Interpret the type as it runs
		fmt.Println(value + 10)
		fmt.Println(value + 10.10)

		
		// const value int = 45
		// fmt.Println(value + 10)
		// fmt.Println(float32(value) + 10.10)

		fmt.Println(first, second)

}