package main

func main() {
	var i int
	//Loop until a condition is satisfied
	for i < 5 {
		println( i)
		i++;
	}

	for j := 0 ; j < 5; j++ {
		println(j)
	}

	
	//Infinite Loop
	var k int
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}
}
