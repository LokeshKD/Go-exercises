package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func simpleMain() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
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
	correct := 0
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, prob.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if prob.a == ans {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
