package main

import "fmt"

func sum(parameters ...int) int {
	total := 0
	for _, num := range parameters {
		total += num
	}

	return total
}

func print(value int) {
	fmt.Println(value)
}
