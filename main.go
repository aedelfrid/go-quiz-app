package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func getProblems() (questions []string, answers []string){
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, _ := reader.ReadAll()

	for _, problem := range data {
		questions = append(questions, problem[0])
		answers = append(answers, problem[1])
	}

	return questions, answers
}

func askQuestion() {

}

func handleAnswer() {

}

func main() {
	questions, answers := getProblems()

	for i, question := range questions {
		fmt.Printf("What is the answer to %s?\n%s\n", question, answers[i])
	}

}
