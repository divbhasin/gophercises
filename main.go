package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"bufio"
)

func main() {
	filePtr := flag.String("filename", "problems.csv", "Name of the CSV file to read")
	flag.Parse()

	score := startGame(*filePtr)
	fmt.Printf("Game ended. Final score: %d", score)
}

func startGame(name string) (score int) {
	f, err := os.Open(name)
	if err != nil {
		return score
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return score
	}

	for _, line := range lines {
		fmt.Printf("Question: %s", line[0])
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == line[1] {
			score++
			fmt.Printf("Correct answer! Score: %d", score)
		}
	}

	return score
}