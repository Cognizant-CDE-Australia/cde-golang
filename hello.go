package main

import (
	"fmt"

	cust "github.com/cdeGolang/golangPackage/customer"
	emp "github.com/cdeGolang/golangPackage/employee"
)

func main() {
	fmt.Println("Welcome to packages")
	fmt.Println("Print customer from pacakge")
	cust.PrintLine()

	cust1 := cust.ReturnCustomer()
	fmt.Println(cust1)

	fmt.Println("Package 2")

	emp.PrintEmployee()
}
