package main

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
