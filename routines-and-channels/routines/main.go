package main

import (
	"fmt"
	"runtime"
	"time"
)

func counter(counterType string) {
	for i := 0; i < 5; i++ {
		fmt.Println(counterType, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// go counter("A")
	// go counter("B")
	// time.Sleep(time.Second * 10)

	runtime.GOMAXPROCS(1)
	fmt.Println("começou")
	go func() {
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Terminou")
}
