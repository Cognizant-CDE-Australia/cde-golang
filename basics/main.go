package main

import (
	"fmt"
	"go-power/com/nike/service/interfaces"
)

type form struct {
	value string
}

func main() {
	//v1 := service.IsPrimeNumber(100)
	//service.IsPrimeNumber(75)

	//log.Println(v1)
	//service.ForEach()

	//numsarray := []int{4, 6, 3, 9}
	//service.TwoSum(numsarray, 9)
	//service.SliceEx1()
	//service.StringConcat()

	//service.Generate()

	var formarray []form

	forms := append(formarray, form{"a"})

	for index1, value1 := range forms {
		fmt.Println(index1, value1)
	}

	v1:= "Hello\n\"vimal\""


	fmt.Println(v1)

	array := []int{1, 2, 4}

	array = append(array, 7)

	for index, value := range array {
		fmt.Println(index, value)
	}


	strconv.p
	interfaces.ShowPower()
}
