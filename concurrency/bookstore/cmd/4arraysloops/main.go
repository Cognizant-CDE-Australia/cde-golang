package main

import "fmt"

const currency = "AUD"

type book struct {
	name     string
	price    int
	author   string
	quantity int
}

func main() {
	emptyBook := new(book)
	//displayBook(emptyBook)

	validBook := &book{name: "Go types", price: 60, quantity: 8, author: "Types"}
	//displayBook(book)
	booksSlice := make([]*book, 0)
	booksSlice = append(booksSlice, emptyBook, validBook)

	otherBooksArray := [2]*book{emptyBook, validBook}
	booksSlice = append(booksSlice, otherBooksArray[:]...)

	for _, book := range booksSlice {
		displayBook(book)
	}

}

func displayBook(book *book) {
	fmt.Printf("The book is %s. it costs %d %s\n", book.name, book.price, currency)
	fmt.Printf("There are  %d books in stock\n", book.quantity)
	fmt.Printf("There author of the book if  %v \n", book.author)
}
