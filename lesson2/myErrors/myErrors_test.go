package myErrors

import (
	"fmt"
)

func Example() {
	err := New("Alarm! An error has occured!")
	fmt.Println(err.Error())
	// Output:
	//Error text: Alarm! An error has occured! | Error time: 2021-12-16 22:00:00.0000 +0300 MSK m=+0.001379126
}
