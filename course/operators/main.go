package main

import (
	"fmt"
	"math"
)

func arithmeticOperators() {
	// +, -, *, /, %, just like you're used to

	// bitwise
	fmt.Println("AND => ", 11&10) // 10
	fmt.Println("OR => ", 11|10)  // 11
	fmt.Println("XOR => ", 11^10) // 01

	fmt.Println(math.Pow(3, 3))
	fmt.Println(math.Ceil(3.4123123123))
	fmt.Println(math.Floor(3.4123123123))
}

func main() {
	arithmeticOperators()
}
