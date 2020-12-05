package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func getBinaryRange(start float64, end float64, partition string) (float64, float64) {
	var lo, hi float64 = 0, 0

	if partition == "lower" {
		lo, hi = start, math.Floor((end+start)/2)
	} else {
		lo, hi = math.Ceil((end+start)/2), end
	}

	return lo, hi
}

func getSeatID(seatCode string, maxRow float64, maxColumn float64) float64 {
	var positions = map[string]string{
		"F": "lower",
		"B": "upper",
		"L": "lower",
		"R": "upper"}

	var seatRowMin float64 = 0
	seatRowMax := maxRow
	for i := 0; i < 7; i++ {
		seatRowMin, seatRowMax = getBinaryRange(seatRowMin, seatRowMax, positions[string(seatCode[i])])
	}

	var seatCloumnMin float64 = 0
	seatColumnMax := maxColumn
	for i := 7; i < 10; i++ {
		seatCloumnMin, seatColumnMax = getBinaryRange(seatCloumnMin, seatColumnMax, positions[string(seatCode[i])])
	}

	return seatRowMin*8 + seatCloumnMin
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var seatIDs []float64
	for scanner.Scan() {
		line := scanner.Text()

		seatID := getSeatID(line, 127, 7)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Float64s(seatIDs)

	// part 1
	fmt.Println(seatIDs[len(seatIDs)-1])

	// part 2
	for i := 0; i < len(seatIDs)-1; i++ {
		current := seatIDs[i]
		next := seatIDs[i+1]

		if next-current > 1 {
			fmt.Println(current + 1)
			break
		}
	}
	file.Close()
}
