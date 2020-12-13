package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func answerMap(groupAnswers string) map[byte]int {
	answers := map[byte]int{}

	for i := 0; i < len(groupAnswers); i++ {
		answers[groupAnswers[i]]++
	}

	return answers
}

// part 1
func countGroupAnswers(groupAnswers string) int {
	return len(answerMap(groupAnswers))
}

// part 2
func countCommonAnswers(groupAnswers string, groupSize int) int {
	answers := answerMap(groupAnswers)

	count := 0
	for _, v := range answers {
		if v == groupSize {
			count++
		}
	}

	return count
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	groupSize := 0
	var groupAnswers []string

	counter := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			counter += countCommonAnswers(strings.Join(groupAnswers, ""), groupSize)
			groupSize = 0
			groupAnswers = nil
		} else {
			groupSize++
			groupAnswers = append(groupAnswers, line)
		}
	}
	counter += countCommonAnswers(strings.Join(groupAnswers, ""), groupSize)

	fmt.Println(counter)
	file.Close()
}
