package main

import "fmt"

const currency = "AUD"

func main() {
	var (
		name             = "Go book"
		price            = 50
		author, quantity = "Google", 5
	)
	displayBook(name, author, quantity, price)
}

func displayBook(name string, author string, quantity int, price int) {
	fmt.Printf("The book is %s. it costs %d %s\n", name, price, currency)
	fmt.Printf("There are  %d books in stock\n", quantity)
	fmt.Printf("There author of the book if  %v \n", author)
}
