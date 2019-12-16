package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filePtr := flag.String("filename", "problems.csv", "Name of the CSV file to read")
	time := flag.Int("limit", 10, "Total time (in seconds) allowed to answer all questions")
	randomize := flag.Bool("randomize", false, "Randomize questions (true or false)")
	flag.Parse()

	score, numQuestions := runGame(*filePtr, *time, *randomize)
	fmt.Printf("Game ended. Final score: %d/%d", score, numQuestions)
}

func runGame(name string, limit int, randomize bool) (score int, numQuestions int) {
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	defer timer.Stop()

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

	if randomize {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(lines), func (i, j int) { lines[i], lines[j] = lines[j], lines[i]})
	}

mainLoop:
	for _, line := range lines {
		fmt.Printf("Question: %s\n", line[0])

		answers := make(chan string)

		go func() {
			fmt.Printf("Answer: ")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.TrimSuffix(text, "\n")
			answers <- text
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime is up!")
			break mainLoop
		case ans := <- answers:
			if ans == line[1] {
				score++
				fmt.Printf("Correct answer! ")
			} else {
				fmt.Printf("Wrong answer! ")
			}
			fmt.Printf("Score: %d\n", score)
		}
	}

	return score, len(lines)
}