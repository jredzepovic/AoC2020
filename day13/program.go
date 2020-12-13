package main

import (
	"fmt"
)

func main() {
	earliestTS := 1003681
	busIDs := []int{23, 37, 431, 13, 17, 19, 409, 41, 29}

	// part 1, brute force
	min := earliestTS * 2
	var minIndex int
	for i, b := range busIDs {
		for j := b; ; j += b {
			if j >= earliestTS {
				if j < min {
					min = j
					minIndex = i
				}
				break
			}
		}
	}

	fmt.Println((min - earliestTS) * busIDs[minIndex])
}
