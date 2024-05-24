package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func getProblems() (data [][]string) {
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, _ = reader.ReadAll()

	return data
}

func serveProblem(problem []string, score int) int {
	// takes a problem and prints question to stdout
	// scans for a response from stdin
	question := problem[0]
	answer := problem[1]

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is %s?\n", question)
	text, _ := reader.ReadString('\n')
	
	isCorrect := handleResponse(text, answer)

	fmt.Printf("Answer is %s.\n", isCorrect)

	if isCorrect == "Correct" {
		score++
	}

	return score

}

func handleResponse(text string, answer string) string{

	// checks stdin against the answer
	if strings.TrimRight(text, "\r\n") != answer {
		return "Incorrect"
	}
	
	return "Correct"
}

func main() {
	problems := getProblems()

	score := 0

	for _, problem := range problems {
		score = serveProblem(problem, score)
	}

	fmt.Printf("Score: %d.\n", score)

}
