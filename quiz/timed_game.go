package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to Parse CSV file %s\n", *csvFileName))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, prob := range problems {
		ansCh := make(chan string)
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.", correct, len(problems))
			return
		case ans := <-ansCh:
			if prob.a == ans {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.", correct, len(problems))
}
