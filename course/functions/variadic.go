package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func printApproved(approved ...string) {
	fmt.Println("Approved List")

	for i, a := range approved {
		fmt.Printf("%d) %s\n", i+1, a)
	}
}
