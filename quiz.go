package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func createProblemList(data [][]string) []Problem {
	var problemList []Problem
	for _, line := range data {
		var rec Problem
		for j, field := range line {
			if j == 0 {
				rec.Question = field
			} else if j == 1 {
				rec.Answer = field
			}
		}
		problemList = append(problemList, rec)

	}
	return problemList
}

func play(QuestionList []Problem) {
	var score int
	timeLimit := flag.Int("Limit", 5, "The time limit for the quiz in seconds")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, rec := range QuestionList {
		fmt.Println(rec.Question)
		answerch := make(chan string)
		go func() {
			//var answer string
			var answer string
			fmt.Scanln(&answer)
			answerch <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Time is up!")
			fmt.Println("Score: %d out of %d", score, len(QuestionList))
			return
		case answer := <-answerch:
			if answer == rec.Answer {
				score++
			}

		}

	}
	fmt.Printf("You scored %d out of %d", score, len(QuestionList))
}

func main() {
	fmt.Println("Hello World")
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(data)

	// convert records to array of structs
	problemList := createProblemList(data)
	//fmt.Println(problemList)
	// print the array
	//fmt.Printf("%+v\n", problemList)

	//play
	play(problemList)
}
