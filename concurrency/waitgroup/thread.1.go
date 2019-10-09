package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	time.Sleep(50 * time.Millisecond)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go say(fmt.Sprintf("thread - %d", i), &wg)
	}
	wg.Wait()
}
