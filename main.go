package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePtr := flag.String("filename", "problems.csv", "Name of the CSV file to read")
	time := flag.Int("timelimit (in seconds)", 30, "Total time allowed to answer all questions")
	flag.Parse()

	score, numQuestions := startGame(*filePtr)
	fmt.Printf("Game ended. Final score: %d/%d", score, numQuestions)
}

func startGame(name string) (score int, numQuestions int) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("Questions file could not be opened")
		return score, 0
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Failed to parse csv file")
		return score, len(lines)
	}

	for _, line := range lines {
		fmt.Printf("Question: %s", line[0])
		fmt.Printf("Answer: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		text = strings.TrimSuffix(text, "\n")

		if text == line[1] {
			score++
			fmt.Println("Correct answer! ")
		} else {
			fmt.Println("Wrong answer! ")
		}

		fmt.Printf("Score: %d\n", score)
	}

	return score, len(lines)
}