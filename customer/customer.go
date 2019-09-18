package customer

import "fmt"

type Customer struct { 
	firstName string
	lastName string
	age int
}

func PrintLine() {
	cust1 := Customer{
		firstName: "John",
		lastName: "Doe",
		age: 45,
	}
	fmt.Println(cust1)
}

func ReturnCustomer() Customer {
	cust1 := Customer{
		firstName: "Smith",
		lastName: "John",
		age: 22,
	}
	return cust1
}
