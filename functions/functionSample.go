package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func calc(a int, b int, f func(int, int) int) int {
	result := f(a, b)
	return result
}

func main() {
	addResult := calc(5, 3, add)
	subResult := calc(6, 3, subtract)
	mulResult := calc(6, 3, mul)

	fmt.Println("5 + 3 = ", addResult)
	fmt.Println("6 - 3 = ", subResult)
	fmt.Println("6 * 3 = ", mulResult)
}
