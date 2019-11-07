package service

import (
	"cde-restservices/com/cde/model"
	"container/list"
	"errors"
	"fmt"
	_ "strconv"
	_ "time"
)
var values = list.New()

func init() {
}



func AddBook(book * model.Book) {
	fmt.Println("Add Book Method Invoked")
	values.PushFront(book)
}


func DeleteBook(id string) {

	for element := values.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(*model.Book))
		book := element.Value.(*model.Book)
		if book.Id == id {
			values.Remove(element)
		}
	}

}

func GetBookById(id string) (*model.Book, error) {

	fmt.Println("Book searched for Id ", id)
	for element := values.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(*model.Book))
		book := element.Value.(*model.Book)
		fmt.Println("Book found for Id ", book.Id)

		if book.Id == id {
			return book, nil
		}

	}
	return nil, errors.New("book with the id not found")
}

func GetAllBooks()   [] *model.Book {

	var slice [] *model.Book


	for element := values.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(*model.Book))
		book := element.Value.(*model.Book)
		slice = append(slice, book)
	}

	fmt.Println("Service ... GetAllBooks")
	//fmt.Println(values)
	//return * values
	return slice
}