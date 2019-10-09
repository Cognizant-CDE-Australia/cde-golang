package main

import (
	"fmt"
	"strconv"
	"time"
)

// Finder
func finder(mine [5]string, oreChannel chan string) {
	fmt.Println("---------start finder: something in the oreChannel-----------")
	for i, item := range mine {
		if item == "ore" {
			fmt.Println("From Finder: processing a ore", item, i)
			oreChannel <- item + strconv.Itoa(i) //send item on oreChannel
		}
	}
	fmt.Println("----------end finder: Ore is found----------------")
}

// Miner
func miner(oreChannel chan string, minedOreChan chan string) {
	fmt.Println("----------start Miner: something in the oreChannel---------")
	for i := 0; i < 3; i++ {
		foundOre := <-oreChannel //read from oreChannel
		fmt.Println("From Miner: Processing  ore ", foundOre)
		minedOreChan <- foundOre //send to minedOreChan
	}
	fmt.Println("----------end miner: Ore is mined------------")
}

// Smelter
func smelter(minedOreChan chan string) {
	fmt.Println("----------start smelter: something in the oreChannel------------")
	for i := 0; i < 3; i++ {
		minedOre := <-minedOreChan //read from minedOreChan
		fmt.Println("From Smelter: Processing Ore ", minedOre)
	}
	fmt.Println("-----------end smelter: Ore is smelted----------------------")
}

func main() {

	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// Finder
	go finder(theMine, oreChannel)
	// Miner
	go miner(oreChannel, minedOreChan)
	// Smelter
	go smelter(minedOreChan)
	<-time.After(time.Second * 5) // Again, you can ignore this
}
