package main

import "fmt"

func arrays() {
	var scores [3]float64
	fmt.Println(scores)

	scores[0], scores[1], scores[2] = 8, 3.1, 4.5
	fmt.Println(scores)

	total := 0.0

	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}

	avg := total / float64(len(scores))
	fmt.Printf("Average: %.2f\n", avg)
}

func looping() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Printf("%d -> %d\n", i, number)
	}
}

func slices() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 11)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3)
	fmt.Println(s, len(s), cap(s))
}

func slicesAsPointers() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	// ?[0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 1 2 3]
	// ? [7 0 0 0 0 0 0 0 0 0] [7 0 0 0 0 0 0 0 0 0 1 2 3]
	fmt.Println(s1, s2)
}

func appending() {
	// array1 := [3]int{1, 2, 3}
	var slice []int

	slice = append(slice, 4, 5, 6, 7, 8, 9)
	fmt.Println(slice)

	slice2 := make([]int, 6)
	copy(slice2, slice[:3])
	fmt.Println(slice2)
}

func main() {
	// arrays()
	// looping()
	// slices()
	// slicesAsPointers()
	appending()
}
