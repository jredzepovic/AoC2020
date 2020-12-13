package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(numbers []int) int {
	var isFound bool = false
	var sum int = 2020
	var result int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				result = numbers[i] * numbers[j]
				isFound = true
				break
			}
		}
		if isFound {
			break
		}
	}

	return result
}

func part2(numbers []int) int {
	var isFound bool = false
	var sum int = 2020
	var result int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == sum {
					result = numbers[i] * numbers[j] * numbers[k]
					isFound = true
					break
				}
			}
			if isFound {
				break
			}
		}
		if isFound {
			break
		}
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to convert input file.")
		}
		numbers = append(numbers, x)
	}

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
	file.Close()
}
