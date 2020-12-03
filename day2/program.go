package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidPassword(policy string, password string) bool {
	var policySplitted = strings.Split(policy, " ")

	var letter = policySplitted[1]

	minCount, err := strconv.Atoi(strings.Split(policySplitted[0], "-")[0])
	if err != nil {
		log.Fatalf("Failed to convert min count.")
	}

	maxCount, err := strconv.Atoi(strings.Split(policySplitted[0], "-")[1])
	if err != nil {
		log.Fatalf("Failed to convert max count.")
	}

	var countInPassword = strings.Count(password, letter)

	if countInPassword < minCount || countInPassword > maxCount {
		return false
	}
	return true
}

func isValidPasswordV2(policy string, password string) bool {
	var policySplitted = strings.Split(policy, " ")

	var letter = policySplitted[1]

	firstPosition, err := strconv.Atoi(strings.Split(policySplitted[0], "-")[0])
	if err != nil {
		log.Fatalf("Failed to convert first position.")
	}

	secondPosition, err := strconv.Atoi(strings.Split(policySplitted[0], "-")[1])
	if err != nil {
		log.Fatalf("Failed to convert second position.")
	}

	if (string(password[firstPosition-1]) == letter) != (string(password[secondPosition-1]) == letter) {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var counter int = 0
	for scanner.Scan() {
		var x = strings.Split(scanner.Text(), ":")

		if isValidPasswordV2(strings.TrimSpace(x[0]), strings.TrimSpace(x[1])) {
			counter++
		}
	}

	fmt.Println(counter)
	file.Close()
}
