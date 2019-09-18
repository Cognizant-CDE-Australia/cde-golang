package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter number : ")
	fmt.Scanf("%d", &num)
	
	if (num%2 == 0) {
		fmt.Println("Its an even number")
	} else {
		fmt.Println("Its an odd number")
	}
}
