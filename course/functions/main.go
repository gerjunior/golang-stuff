package main

import "fmt"

func init() {
	fmt.Println("Starting....")
}

func main() {
	// result := sum(5, 3, 10, 20)
	// print(result)

	// fmt.Printf("Average: %.2f\n", average(8.8, 7.9, 1.2))

	// approved := []string{"John", "Jane", "Ralph"}
	// printApproved(approved...)

	// memorized := closure()
	// memorized()
	// // + 1
	// memorized()
	// result, _ := factorial(5)
	// fmt.Println(result)

	// _, err := factorial(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(factorial(5))

	n := 1

	increment(&n)
	increment(&n)
	fmt.Println(n)
}
