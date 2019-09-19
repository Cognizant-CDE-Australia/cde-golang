package service

import "log"

func SliceEx1() {
	var array []string

	array = append(array, "Usa")
	array = append(array, "Canada")

	log.Println(array)
}
