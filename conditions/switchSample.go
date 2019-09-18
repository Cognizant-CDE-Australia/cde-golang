package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter number : ")
	fmt.Scanf("%d", &num)
	
	switch num%2 {
		case 0:
			fmt.Println("Input is a even number")
			break;
		case 1:
			fmt.Println("Input is a odd number")
			break;
		default:
			fmt.Println("Input is not a number")
			break;
	}
}