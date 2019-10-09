package pkg

type book struct {
	name     string
	price    int
	author   string
	quantity int
}

// Newbook book
//is a function to create a new book
func NewBook(name string, author string, quantity int, price int) book {
	book := &book{name: name, price: price, quantity: quantity, author: author}
	return book
}
