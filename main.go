package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func startTimer(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	fmt.Println("Times up!")
	os.Exit(0)
}

func main () {
	q := Quiz{}
	reader := bufio.NewReader(os.Stdin)
	// flags
	problemFile := flag.String("pf", "problems.csv", "A csv containing problems")
	quizTime := flag.Int("t", 10 , "Timer duration for quiz")
	// parse command line
	flag.Parse()
	go startTimer(*quizTime)
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
	// loop through csv and build quiz
	for _, line := range lines {
		q.AddProblem(Problem{line[0],line[1]})
	}
	// loop through questions and ask user to solve problems
	for _, p := range q.Problems {
		fmt.Printf("Question: What is %s?\n", p.Question)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text,"\n","", -1)
		if strings.Compare(p.Answer, text) == 0 {
			fmt.Println("You're right!")
		} else {
			fmt.Println("Wrong answer.")
		}
	}

}
