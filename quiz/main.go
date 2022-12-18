package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "./problems.csv", "a csv file in the format of 'question,answer'")
	duration := flag.Duration("d", time.Second*30, "limit time to answer the questions. If the time runs out, the program stops and the score shows up.")
	flag.Parse()

	timer := time.NewTimer(*duration)
	defer timer.Stop()

	correct := 0
	problems := parseProblems(*csvFileName)

	showScore := func() {
		fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
	}

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			showScore()
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	showScore()
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

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
