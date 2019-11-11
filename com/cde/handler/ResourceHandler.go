package handler

import (
	"cde-restservices/com/cde/model"
	"cde-restservices/com/cde/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add book Invoked")

	fmt.Fprint(w, "fields=" + r.URL.Query().Get("fields") + "\n")

	paramArray := r.URL.Query()

	for key, value := range paramArray {
		fmt.Println(key, value)
	}

	fmt.Println(r.URL.Query().Get("g1"))

	//vars := mux.Vars(r)
	//title := vars["Title"]
	//id := vars["Id"]
	//
	//fmt.Println("The incoming vars are", title, id )

	reqBody, _ := ioutil.ReadAll(r.Body)
	var book model.Book
	json.Unmarshal(reqBody, &book)
	fmt.Println(book)
	service.AddBook(&book)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update book Invoked")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var book model.Book
	json.Unmarshal(reqBody, &book)
	fmt.Println(book)
	service.UpdateBookById(&book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete book Invoked")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book model.Book
	json.Unmarshal(reqBody, &book)
	fmt.Println(book)
	service.DeleteBook(book.Id)
}


func BookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Book By Id Invoked")
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var book model.Book
	//json.Unmarshal(reqBody, &book)
	w.Header().Set("Content-Type", "application/json")


	vars := mux.Vars(r)
	id := vars["id"]

	foundBook, err := service.GetBookById(id)
	if err != nil {
		http.Error(w, err.Error(), 404)

	} else {
		json.NewEncoder(w).Encode(foundBook)
	}


}


func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books Invoked")


	foundBooks := service.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")

	//if (foundBook != nil)
	json.NewEncoder(w).Encode(foundBooks)


}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
