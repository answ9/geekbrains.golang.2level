package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

//Если закомментировать runtime.Gosched(), убрать трассировку и запустить на одном ядре (GOMAXPROCS=1),
//то программа выводет 5 раз только "hello"
//В текущем состоянии в консоль поочередно будет выводиться "hello" и "world".
//Принудительный вызов планировщика помогает получить время выполнения другим потокам программы
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	go say("world")
	say("hello")
}
