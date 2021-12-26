package main

import (
	"fmt"
	"lesson2/myErrors"
	"log"
)

func main() {
	err := myFuncWithRecoveredPanic()
	if err != nil {
		log.Println(err)
	}
}

func myFuncWithRecoveredPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = myErrors.New(fmt.Sprintf("recovered %v", v))
		}
	}()

	arr := []uint{0, 1, 2, 3}
	for i := 0; i <= 5; i++ {
		fmt.Printf("index: %d, value: %d\n", i, arr[i])
	}

	return
}
