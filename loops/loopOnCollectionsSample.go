package main

func main() {

	stringSlice := []string{"Hello", "John", "Smith"}

	//Legacy way of iterating an array / slice
	for i := 0; i < len(stringSlice); i ++ {
		println(stringSlice[i])
	}

	//New way of iterating in go
	for i, v := range stringSlice {
		println(i, v)
	}

	employee := map[int]string {1 : "John", 2: "Smith", 3: "Bravo"}
	for i, v := range employee {
		println(i, v)
	}
}
