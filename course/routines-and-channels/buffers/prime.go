package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int, c chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime
				fmt.Printf("%d was stored in channel buffer\n", prime)
				start = prime + 1
				break
			}
		}
	}
	close(c)
}
