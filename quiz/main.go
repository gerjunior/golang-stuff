package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "./problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	problems := parseProblems(*csvFileName)

	correct := 0
	for i, p := range problems {
		answer := scanAnswer(i+1, p.question)

		if answer == p.answer {
			correct += 1
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseProblems(csvFileName string) []problem {
	file, err := os.Open(csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	return problems
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func scanAnswer(number int, question string) string {
	fmt.Printf("Problem #%d: %s = \n", number, question)
	var answer string
	fmt.Scanf("%s\n", &answer)
	return answer
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
