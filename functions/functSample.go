package main

import "fmt"

func main() { 
	fmt.Println("Hello function")

	helloFunction()

	defer deferFunction()
	
	paramFunction("John", 45)

	result1 := addTwoNumbers(1, 5)
	fmt.Println(result1)

	result2, result3 := divideTwoNumbers(10, 3)
	fmt.Println(result2, result3)

	//Omitting some variable values
	_, result4 := divideTwoNumbers(10, 3)
	fmt.Println(result4)

	result4, result5 := mathOperation(2, 6)
	fmt.Println(result4, result5)

	result6 := getFactorialNumber(4)
	fmt.Println(result6)
	
}

func helloFunction() {
	fmt.Println("Hello function called....")
}

func paramFunction(name string, age int) {
	fmt.Printf("Hello %s age %d\n", name, age)
}

func addTwoNumbers(a, b int) int64 {
	return int64(a + b)
}

//multiple return values
func divideTwoNumbers(a, b int) (int, int) {
	return a/b, a%b
}

func mathOperation(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}

//Recursive function
func getFactorialNumber(num int) int {
	if num > 1 {
		return num * getFactorialNumber(num-1)
	} else {
		return 1 // 1! == 1
	}
}

func deferFunction(){
	fmt.Println("End of functions chapter")
}