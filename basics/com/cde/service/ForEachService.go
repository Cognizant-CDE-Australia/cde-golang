package service

import "log"

func ForEach() {

	countries := []string{"Usa", "Canada", "Germany", "Switzerland"}

	for index, value := range countries {
		log.Println(index, value)
	}
}
