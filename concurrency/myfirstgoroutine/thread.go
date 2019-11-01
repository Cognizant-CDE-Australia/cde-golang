package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
	time.Sleep(3 * time.Second)
}

func main() {

	for i := 0; i < 5; i++ {
		go say(fmt.Sprintf("world %d", i))
	}

	// time.Sleep(20 * time.Second)
}
