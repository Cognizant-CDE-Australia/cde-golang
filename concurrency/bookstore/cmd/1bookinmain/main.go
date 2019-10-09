package main

import "fmt"

const currency = "AUD"

func main() {
	var name string
	name = "Go book"
	author := "Google"
	price, quantity := 50, 5
	fmt.Printf("The book is %s. it costs %d %s\n", name, price, currency)
	fmt.Printf("There are  %d books in stock\n", quantity)
	fmt.Printf("There author of the book if  %v \n", author)
}
