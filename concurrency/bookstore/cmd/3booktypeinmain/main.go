package main

import "fmt"

const currency = "AUD"

func main() {
	emptyBook := new(book)
	displayBook(emptyBook)

	book := &book{name: "Go types", price: 60, quantity: 8, author: "Types"}
	displayBook(book)
}

func displayBook(book *book) {
	fmt.Printf("The book is %s. it costs %d %s\n", book.name, book.price, currency)
	fmt.Printf("There are  %d books in stock\n", book.quantity)
	fmt.Printf("There author of the book if  %v \n", book.author)
}

type book struct {
	name     string
	price    int
	author   string
	quantity int
}
