package main

import (
	"fmt"
)

func runGame(steps int, initialNumbers []int) int {
	round := 1
	game := map[int][]int{}

	for _, i := range initialNumbers {
		game[i] = append(game[i], round)
		round++
	}

	lastNumber := initialNumbers[len(initialNumbers)-1]
	for {
		if round > steps {
			break
		}
		if len(game[lastNumber]) == 1 {
			lastNumber = 0
			game[lastNumber] = append(game[lastNumber], round)
		} else {
			lastNumber = game[lastNumber][len(game[lastNumber])-1] - game[lastNumber][len(game[lastNumber])-2]
			game[lastNumber] = append(game[lastNumber], round)
		}
		round++
	}

	return lastNumber
}

func main() {
	// part 1
	fmt.Println(runGame(2020, []int{2, 15, 0, 9, 1, 20}))

	// part 2
	fmt.Println(runGame(30000000, []int{2, 15, 0, 9, 1, 20}))
}
