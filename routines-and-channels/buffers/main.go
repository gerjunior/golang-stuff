package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	fmt.Println("Buffering: 1")
	ch <- 2
	fmt.Println("Buffering: 2")
	ch <- 3
	fmt.Println("Buffering: 3")
	ch <- 4
	fmt.Println("Buffering: 4")
	ch <- 5
	fmt.Println("Buffering: 5")
	ch <- 6
	fmt.Println("Buffering: 6")
	ch <- 7
	fmt.Println("Buffering: 7")
}

// func main() {
// 	ch := make(chan int, 3)
// 	go routine(ch)

// 	time.Sleep(time.Second * 2)
// 	// ? after 2 seconds the function below will read the first chunk of the buffer,
// 	// ? thus making the channel read one more chunk from the routine
// 	fmt.Println("Reading: ", <-ch)
// 	time.Sleep(time.Second * 2)
// }

// func main() {
// 	ch := make(chan int, 3)

// 	routine(ch)

// 	for i := range ch {
// 		fmt.Println(i)
// 	}
// }

func main() {
	ch := make(chan int, 99)
	amountPrimes := 100
	go primes(amountPrimes, ch)

	for i := 0; i < amountPrimes; i++ {
		fmt.Printf("Reading: %d\n", <-ch)
		time.Sleep(time.Second)
	}
}
