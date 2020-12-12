package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkPairSum(number int, numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if (numbers[i] != numbers[j]) && (numbers[i]+numbers[j] == number) {
				return true
			}
		}
	}
	return false
}

func sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func min(numbers []int) int {
	m := numbers[0]
	for _, n := range numbers {
		if n < m {
			m = n
		}
	}
	return m
}

func max(numbers []int) int {
	m := numbers[0]
	for _, n := range numbers {
		if n > m {
			m = n
		}
	}
	return m
}

func encryptionWeakness(number int, numbers []int) int {
	for s := 2; s < len(numbers); s++ {
		for i := len(numbers) - 2; i+1 >= s; i-- {
			subslice := numbers[i-s+1 : i+1]

			if number == sum(subslice) {
				return min(subslice) + max(subslice)
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	preamble := 25
	numbers := []int{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
		i++

		if i > preamble && !checkPairSum(n, numbers[i-preamble-1:i]) {
			// part 1
			fmt.Println(n)

			// part 2
			fmt.Println(encryptionWeakness(n, numbers))
			break
		}
	}
	file.Close()
}
