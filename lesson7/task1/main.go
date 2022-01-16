package main

import (
	"fmt"
	"log"
	"task1/funcs"
	"task1/persons"
)

func main() {
	m := map[string]interface{}{
		"Name":        "Vladimir",
		"Age":         31,
		"Married":     false,
		"Temperature": 36.6,
	}

	person := persons.Person{}
	err := funcs.ChangeStructField(&person, m)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(person)
}
