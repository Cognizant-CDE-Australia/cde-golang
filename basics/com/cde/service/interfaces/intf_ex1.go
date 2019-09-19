package interfaces

import "fmt"

type Power interface {
	getPower() int
}

type Prado struct {
}

type Pajero struct {
}

func (Pajero) getPower() int {
	return 500
}

func (Prado) getPower() int {
	return 600
}

func ShowPower() {
	prado := Prado{}
	prado.getPower()
	InterfacePower(prado)
}

func InterfacePower(p Power) {
	fmt.Println(p.getPower())
}
