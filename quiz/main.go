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
	flags := getFlags()
	questions := parseCsv(flags.csvFileName)

	fmt.Printf("You have %v to answer %v questions. Are you ready?\n", flags.duration, len(questions))
	countdown(5)

	timer := time.NewTimer(flags.duration)
	defer timer.Stop()

	correct := 0
	for i, record := range questions {
		problem := Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}

		fmt.Printf("Question #%d: %s = ", i+1, problem.question)
		inputCh := make(chan string)

		go func() {
			var input string
			fmt.Scanf("%s\n", &input)
			inputCh <- input
		}()

		select {
		case res := <-inputCh:
			if res == problem.answer {
				correct++
			}
		case <-timer.C:
			fmt.Printf("\nTime's up! You scored %d out of %d!\n", correct, len(questions))
			return
		}

	}

	fmt.Printf("You scored %d out of %d!\n", correct, len(questions))
}

type Flags struct {
	csvFileName string
	duration    time.Duration
}

type Problem struct {
	question string
	answer   string
}

func getFlags() Flags {
	csvFileName := flag.String("csv", "./problems.csv", "csv file")
	duration := flag.Duration("duration", time.Second*15, "duration timeout to answer all questions")
	flag.Parse()

	return Flags{
		csvFileName: *csvFileName,
		duration:    *duration,
	}
}

func parseCsv(csvFileName string) [][]string {
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		exit("There was an error while trying to read the CSV file")
	}

	questions, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		exit("There was an error while parsing your csv file")
	}

	return questions
}

func countdown(seconds int) {
	ticker := time.NewTicker(time.Second)

	for i := seconds; i >= 0; i-- {
		select {
		case <-ticker.C:
			if i == 0 {
				fmt.Print(" GO! \n\n")
				ticker.Stop()
			} else {
				fmt.Printf(" %d ", i)
			}
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
