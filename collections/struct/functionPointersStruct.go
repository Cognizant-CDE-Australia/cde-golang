package main

import "fmt"

type Student struct {
	fname, city string
	postcode    int
}

func main() {
	var student1 = Student{"John", "Melbourne", 3000}
	var student2 = Student{"Steve", "Sydney", 2000}
	var student3 = Student{"Mike", "Perth", 4000}

	fmt.Println("student1 Before ", student1)
	fmt.Println("student2 Before ", student2)
	fmt.Println("student3 Before ", student3)

	updateValue(&student3)
	updateValue(&student2)
	updateValue(&student1)

	fmt.Println("student3 After ", student3)
	fmt.Println("student2 After ", student2)
	fmt.Println("student1 After ", student1)
}

func updateValue(student *Student) {
	if student.fname == "Mike" {
		student.fname = "Micheal"
	} else {
		student.postcode = 0
	}
}
