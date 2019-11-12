package main

import (
	"fmt"
	"time"
)

func say(s string, messages chan string) {
	time.Sleep(1 * time.Second)
	messages <- s

}

func main() {

	messages := make(chan string, 5)

	for i := 0; i < 5; i++ {
		go say(fmt.Sprintf("world %d", i), messages)
	}

	i := 0
	for k := range messages {
		i++
		fmt.Println(k)

		if i == 5 {
			close(messages)
		}
	}
}
