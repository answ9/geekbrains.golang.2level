package main

import (
	"fmt"
	"github.com/answ9/test_modules"
	"log"
)

func main() {
	//использую v1
	//fmt.Println(test_modules.Hi("Vladimir"))

	//использую v2
	str, err := test_modules.Hi("Vladimir", "fr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
