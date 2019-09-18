package main

import "fmt"

func main() {
	m := map[int]string{1: "One", 2: "Two", 3: "Three"}
	fmt.Println(m)

	fmt.Println(m[2])

	m[3] = "Thrity three"
	fmt.Println(m[3])

	fmt.Println(m)

	delete(m, 3)

	fmt.Println(m)

	//Initializing a map using make function
	var m1 = make(map[string]int)

	fmt.Println(m1)

	if m1 == nil {
		fmt.Println("m1 is nil")
	} else {
		fmt.Println("m1 is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m1["one hundred"] = 100
	fmt.Println(m1)

	//Other ways of initializing a map

	var m2 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}

	fmt.Println(m2)

	fmt.Println(m2["four"])
}