package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./problems.csv")
	if err != nil {
		panic(err)
	}

	stringFile := string(file)
	questions := strings.Split(stringFile, "\n")

	correct := 0
	for _, line := range questions {
		res := strings.Split(line, ",")
		question := res[0]
		a := res[1]

		answer, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s = ", question)

		var response string
		fmt.Scanln(&response)

		responseInt, err := strconv.Atoi(response)
		if err != nil {
			continue
		}

		if answer == responseInt {
			correct++
		}
	}

	fmt.Printf("You got %v of %v correctly.\n", correct, len(questions))
}
