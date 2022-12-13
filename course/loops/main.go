package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// traditional for loop
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	// while (true)
	for {
		fmt.Println("Forever loop")
		time.Sleep(time.Second)
	}
}
