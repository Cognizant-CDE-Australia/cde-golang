package service

import "fmt"

func IsPrimeNumber(number int) bool {

	fmt.Println("Function Invoked with number :", number)
	if number%2 == 0 {
		return true
	} else {
		return false
	}

}
