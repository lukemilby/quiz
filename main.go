package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)
type Problem struct {
	Question string
	Answer string
}

type Quiz struct {
	Problems []Problem
}

func (q *Quiz) AddProblem(problem Problem) []Problem {
	q.Problems = append(q.Problems, problem)
	return q.Problems
}

func main () {
	q := Quiz{}
	problemFile := flag.String("problemFile", "problems.csv", "A csv containing problems")
	// parse command line
	flag.Parse()

	// Open file
	f, err := os.Open(*problemFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// read csv to lines
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	// loop through csv
	for _, line := range lines {
		q.AddProblem(Problem{line[0],line[1]})
	}
	fmt.Println(len(q.Problems))
}
