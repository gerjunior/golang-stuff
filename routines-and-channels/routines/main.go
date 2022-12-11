package main

import (
	"fmt"
	"time"
)

// func counter(counterType string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(counterType, i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func main() {
// 	// go counter("A")
// 	// go counter("B")
// 	// time.Sleep(time.Second * 10)

// 	runtime.GOMAXPROCS(1)
// 	fmt.Println("começou")
// 	go func() {
// 		for {

// 		}
// 	}()

// 	time.Sleep(time.Second)
// 	fmt.Println("Terminou")
// }

// func main() {
// 	fmt.Println(runtime.NumCPU())
// }

func talk(person, text string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

// func main() {
// 	talk("Maria", "Why don't you talk with me?", 3)
// 	talk("João", "I can only talk after you!", 1)
// }

func main() {
	go talk("Maria", "Hi...", 10)
	go talk("João", "Hello...", 10)

	time.Sleep(time.Second * 5)
}
