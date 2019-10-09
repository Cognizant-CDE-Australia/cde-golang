package main

import (
	"fmt"
	"strconv"
)

// Finder
func finder(mine [5]string) []string {
	var foundore []string
	fmt.Println("start Finder: found a ore", foundore)
	for i, item := range mine {
		if item == "ore" {
			fmt.Println("From Finder: processing ore", item, i)
			foundore = append(foundore, item+strconv.Itoa(i))
		}
	}

	fmt.Println("end Finder: found a ore", foundore)
	return foundore
}

// Ore Breaker
func miner(foundOre []string) []string {
	fmt.Println("start Miner: something in the oreChannel", foundOre)
	for ore := range foundOre {
		fmt.Println("From Miner: Processing ", foundOre[ore])

	}
	fmt.Println("end Miner: something in the oreChannel", foundOre)
	return foundOre

}

// Smelter
func smelter(minedOre []string) []string {
	fmt.Println("start Smelter: Ore is smelted", minedOre)
	for ore := range minedOre {
		fmt.Println("From Smelter: Processing Ore", minedOre[ore])
	}
	fmt.Println("end Smelter: Ore is smelted", minedOre)
	return minedOre
}

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	foundOre := finder(theMine)
	fmt.Println()
	minedOre := miner(foundOre)
	fmt.Println()
	smelter(minedOre)
	fmt.Println()
}
