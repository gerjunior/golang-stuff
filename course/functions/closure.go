package main

import "fmt"

func init() {
	fmt.Println("Starting closure file...")
}

func closure() func() {
	x := 10
	var function = func() {
		fmt.Println(x)
		x++
	}

	return function
}
