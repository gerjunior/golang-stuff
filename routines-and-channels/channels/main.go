package main

import (
	"fmt"
	"time"
)

// func main() {
// thread 1
// thread 1 <-> thread 2
// hello := make(chan string)

// thread 2
// 	go func() {
// 		hello <- "Hello World"
// 	}()

// 	result := <-hello
// 	fmt.Println(result)
// }

// func main() {
// 	forever := make(chan string)

// 	go func() {
// 		x := true
// 		for {
// 			if x == true {
// 				continue
// 			}
// 		}
// 	}()

// 	fmt.Println("Waiting forever...")
// 	<-forever
// }

// func main() {
// 	hello := make(chan string)

// 	go func() {
// 		hello <- "hello, world!"
// 	}()

// 	select {
// 	case x := <-hello:
// 		fmt.Println(x)

// 	default:
// 		fmt.Println("default")
// 	}

// 	time.Sleep(time.Second)
// }

// func main() {
// 	queue := make(chan int)

// 	go func() {
// 		i := 0
// 		for {
// 			time.Sleep(time.Second)
// 			queue <- i
// 			i++
// 		}
// 	}()

// 	for x := range queue {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	ch := make(chan int, 1)

// 	ch <- 1 // sending data to a channel
// 	<-ch    // receiving data from the channel

// 	ch <- 2

// 	fmt.Println(<-ch)
// }

func run(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

// ? just like nodeJS iterators and generators!
// ? yield is similar to -> (send data)
// ? <- is similar to generator.next (read data)
func main() {
	c := make(chan int)
	go run(2, c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
