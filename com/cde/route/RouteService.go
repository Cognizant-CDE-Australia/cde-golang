package route

import (
	"cde-restservices/com/cde/handler"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "strconv"
)

func RegisterRoutes() {
	fmt.Println("Starting the server at 7070 this is v2.................")

	r := mux.NewRouter()

	r.HandleFunc("/", handler.HealthHandler).Methods("GET")
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	r.HandleFunc("/book", handler.AddBook).Methods("POST")
	r.HandleFunc("/book/{id}", handler.BookById).Methods("GET")
	r.HandleFunc("/book", handler.GetAllBooks).Methods("GET")
	r.HandleFunc("/book", handler.DeleteBook).Methods("DELETE")

	fmt.Println("Starting the server at 7070 this is v2")
	_ = http.ListenAndServe(":7070", r)
}



