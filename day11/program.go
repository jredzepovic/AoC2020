package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type direction struct {
	a int
	b int
}

func getNeighbors(i int, j int, state [][]string, step float64) []string {
	neighbors := []string{}
	directions := []direction{
		direction{a: -1, b: 0},
		direction{a: -1, b: 1},
		direction{a: 0, b: 1},
		direction{a: 1, b: 1},
		direction{a: 1, b: 0},
		direction{a: 1, b: -1},
		direction{a: 0, b: -1},
		direction{a: -1, b: -1}}

	for _, d := range directions {
		s := float64(0)
		x := i
		y := j
		for {
			s++
			x = x + d.a
			y = y + d.b
			if x >= 0 && x < len(state) && y >= 0 && y < len(state[0]) {
				if state[x][y] != "." {
					neighbors = append(neighbors, state[x][y])
					break
				}
			} else {
				break
			}
			if s >= step {
				break
			}
		}
	}
	return neighbors
}

func updatePosition(i int, j int, state [][]string) (bool, string) {
	// for part 1, set step to float64(1)
	step := math.Inf(0)
	// for part 1, put 4
	occupied := 5

	if state[i][j] == "." {
		return false, "."
	}

	nCount := map[string]int{}
	neighbors := getNeighbors(i, j, state, step)
	for _, n := range neighbors {
		nCount[n]++
	}

	if state[i][j] == "L" {
		if nCount["#"] == 0 {
			return true, "#"
		}
	} else {
		if nCount["#"] >= occupied {
			return true, "L"
		}
	}
	return false, state[i][j]
}

func updateState(state [][]string) (bool, [][]string) {
	updated := false
	m := [][]string{}
	for range state {
		m = append(m, []string{})
	}

	for i, row := range state {
		for j := range row {
			upd, new := updatePosition(i, j, state)

			m[i] = append(m[i], new)
			if upd {
				updated = true
			}
		}
	}

	return updated, m
}

func countOccupiedPositions(state [][]string) int {
	count := 0

	for _, row := range state {
		for _, col := range row {
			if col == "#" {
				count++
			}
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

	state := [][]string{}
	for scanner.Scan() {
		state = append(state, strings.Split(scanner.Text(), ""))
	}

	updated := false
	for {
		updated, state = updateState(state)
		if !updated {
			break
		}
	}

	fmt.Println(countOccupiedPositions(state))
	file.Close()
}
