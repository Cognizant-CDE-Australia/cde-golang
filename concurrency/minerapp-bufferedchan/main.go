package main

import (
	"fmt"
	"strconv"
	"time"
)

// Finder
func finder(mine []string, oreChannel chan string) {
	fmt.Println("---------start finder: something in the oreChannel-----------")
	for i, item := range mine {
		if item == "ore" {
			fmt.Println("From Finder: processing a ore", item, i)
			oreChannel <- item + strconv.Itoa(i) //send item on oreChannel
		}
	}
	fmt.Println("----------end finder: Ore is found----------------")
	close(oreChannel)
}

// Miner
func miner(oreChannel chan string, minedOreChan chan string) {
	fmt.Println("----------start Miner: something in the oreChannel---------")
	for v := range oreChannel {
		fmt.Println("From Miner: Processing  ore ", v)
		minedOreChan <- v //send to minedOreChan
	}
	fmt.Println("----------end miner: Ore is mined------------")
	close(minedOreChan)
}

// Smelter
func smelter(minedOreChan chan string) {
	fmt.Println("----------start smelter: something in the oreChannel------------")
	for V := range minedOreChan {
		fmt.Println("From Smelter: Processing Ore ", V)
	}
	fmt.Println("-----------end smelter: Ore is smelted----------------------")
}

func main() {

	theMine := []string{"rock", "ore", "ore", "rock", "ore", "tt", "ore", "ore"}
	oreChannel := make(chan string, 2)
	minedOreChan := make(chan string, 4)
	// Finder
	go finder(theMine, oreChannel)
	// Miner
	go miner(oreChannel, minedOreChan)
	// Smelter
	go smelter(minedOreChan)
	<-time.After(time.Second * 5) // Again, you can ignore this
}
