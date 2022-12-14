package main

import "fmt"

func init() {
	fmt.Println("Starting functions file...")
}

func sum(parameters ...int) int {
	total := 0
	for _, num := range parameters {
		total += num
	}

	return total
}

var anonymousSum = func(a, b int) int {
	return a + b
}

func testing() int {
	anonymSum := func(a, b int) int {
		return a + b
	}

	return anonymSum(1, 2)
}

func returningAFunction(namedFunction func(int, int) int, p1, p2 int) int {
	return namedFunction(p1, p2)
}

func exec() int {
	function := func(a, b int) int {
		return a + b
	}
	result := returningAFunction(function, 1, 2)
	return result
}

func print(value int) {
	fmt.Println(value)
}
