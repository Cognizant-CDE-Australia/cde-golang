package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func main() {
	var p1 person
	fmt.Println(p1)

	p1.firstName = "John"
	p1.lastName = "Smith"
	p1.age = 25

	fmt.Println(p1)

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)

	//Other way of assigning value to a type struct 
	p2 := person{firstName: "Brown", lastName: "Fox", age: 10,}
	fmt.Println(p2)

	//Assigning struct value to another struct just like variable initialization
	p3 := p1
	p3.age = 20

	fmt.Println(p3)
}