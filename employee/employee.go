package employee

import "fmt"

// Employee Struct
type Employee struct {
	firstName string
	lastName  string
	company   string
}

// Print Employeeßß
func PrintEmployee() {
	emp1 := Employee{
		firstName: "Hari",
		lastName:  "Dhanakoti",
		company:   "Cognizant",
	}
	fmt.Println("Employee ", emp1)
}
