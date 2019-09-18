package main

import (
	"fmt"
	cust "github.com/haridhanakoti/golangPackage/customer"
)

func main() {
	fmt.Println("Welcome to packages")
	fmt.Println("Print customer from pacakge")
	cust.PrintLine()

	cust1 := cust.ReturnCustomer()
	fmt.Println(cust1)
}