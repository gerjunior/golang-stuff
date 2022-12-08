package main

import (
	"fmt"
	"strconv"
)

func main() {

	for {
		fmt.Print("Enter a number: ")
		var input string
		fmt.Scan(&input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("oops! I was not able to read your number. Why don't you try again?")
			continue
		}

		answer := ""

		if number%3 == 0 {
			answer += "Fizz"
		}

		if number%5 == 0 {
			answer += "Buzz"
		}

		if len(answer) == 0 {
			answer = input
		}

		fmt.Println("=>", answer)
	}
}
