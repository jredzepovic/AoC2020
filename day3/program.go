package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countTrees(x int, y int, area []string) int {
	counter := 0
	i := 0
	for j := y; j < len(area); j += y {
		i = (i + x) % len(area[j])

		if string(area[j][i]) == "#" {
			counter++
		}
	}
	return counter
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var area []string
	for scanner.Scan() {
		area = append(area, scanner.Text())
	}

	fmt.Println(countTrees(3, 1, area))

	fmt.Println(countTrees(1, 1, area) *
		countTrees(3, 1, area) *
		countTrees(5, 1, area) *
		countTrees(7, 1, area) *
		countTrees(1, 2, area))
	file.Close()
}
