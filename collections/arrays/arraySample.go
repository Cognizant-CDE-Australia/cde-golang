package main

import "fmt"

func main() {
	var intData [3]int

	intData[0] = 100
	intData[1] = 200
	intData[2] = 300

	fmt.Println(intData)
	fmt.Println(intData[1])

	newData := [3]int{101, 201, 301}
	fmt.Println(newData)


	var stringData [3]string

	stringData[0] = "John"
	stringData[1] = "Doe"
	stringData[2] = "Smith"

	fmt.Println(stringData)
	fmt.Println(stringData[1])

	newStrData := [3]string{"John", "Doe", "Smith"}
	fmt.Println(newStrData)
}