package service

import "fmt"

type Model struct {
	model string
	year  int
}

type Mitsubhishi struct {
	m Model
}

type Toyota struct {
	m Model
}

func Generate() {
	m1 := Model{
		"Pajero",
		2020,
	}
	m2 := Model{
		"Prado",
		2020,
	}

	mitsubhishi1 := Mitsubhishi{
		m1,
	}
	toyota1 := Toyota{
		m2,
	}

	fmt.Println(mitsubhishi1)
	fmt.Println(toyota1)
}
