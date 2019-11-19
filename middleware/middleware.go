package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Adapter wraps an http.Handler with additional functionality.
type Adapter func(http.Handler, *interface{}) http.Handler

// Adapt h with all specified adapters.
func Adapt(handler interface{}, adapters ...Adapter) (h http.Handler) {
	var response interface{}
	switch handler := handler.(type) {
	case http.Handler:
		h = handler
	case func(http.ResponseWriter, *http.Request):
		h = http.HandlerFunc(handler)
	case func(*http.Request) interface{}:
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response = handler(r)
		})
	default:
		log.Fatal("Invalid Adapt Handler", handler)
	}

	for _, adapter := range adapters {
		h = adapter(h, &response)
	}

	return h
}

// Logging test here
func Logging(l *log.Logger) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("logger", r.Method, r.URL.Path)
			defer func() { l.Printf("logger response %v\n", response) }()
			h.ServeHTTP(w, r)
		})
	}
}

// API adapter
func API() Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			h.ServeHTTP(w, r)

			var x interface{}
			switch v := (*response).(type) {
			case error:
				x = v.Error()
			case string:
				x = v
			default:
				x = "Couldn't figure it out"
			}

			json.NewEncoder(w).Encode(x)
		})
	}
}

// Simplest handler we could write
func apiHandler(r *http.Request) interface{} {
	if r.URL.Path != "/api" {
		return errors.New("API Handler Error")
	}

	return "Success!"
}


// Simplest handler we could write
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	router := http.NewServeMux()
	router.Handle("/api", Adapt(apiHandler, API(), Logging(logger)))
	router.Handle("/invalid", Adapt(apiHandler, API(), Logging(logger)))
	router.HandleFunc("/", helloHandler)

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}