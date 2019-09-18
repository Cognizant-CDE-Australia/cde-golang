package main

import "fmt"

func main() {

	intData := [5]int{100, 200, 300, 400, 500}

	slice := intData[:]
	fmt.Println(intData, slice)

	intData[1] = 101

	slice[2] = 201

	fmt.Println(intData, slice)

	// Slice without an array
	strSlice := []string{"One", "Two", "Three"}
	fmt.Println(strSlice)

	strSlice = append(strSlice, "Four")
	fmt.Println(strSlice)

	strSlice = append(strSlice, "Five", "Six", "Seven")
	fmt.Println(strSlice)

	//Traversing the slice with the index
	strSlice2 := strSlice[1:]
	strSlice3 := strSlice[:2]
	strSlice4 := strSlice[3:5]
	fmt.Println(strSlice2, strSlice3, strSlice4)


}