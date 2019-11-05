package main

import (
	"os"
)

func main() {
	//Create panic with a message
	// panic("Error / exception message")

	_, err := os.Create("c:/tmp/file")
	if err != nil {
		panic(err)
	}

	var m map[string]int
	if m == nil {
		println("m is nil")
		panic("Map is empty")
	}
}
