package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'") // This is a pointer to string
	timeLimit := flag.Int("limit", 30, "the time limit to the quiz in seconds")
	flag.Parse()

	_ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s!", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)

		// None blocking
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// Blocking. Dont put defaul block
		select {
		case <-timer.C:
			fmt.Printf("You scored %d of %d\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

		}
	}
	fmt.Printf("You scored %d of %d\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	// We already  knew the expected length so we should not use append
	// Which need to resize the slice everytime new element is added
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1])}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
