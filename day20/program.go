package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	ID          int
	borderTop   string
	borderBot   string
	borderLeft  string
	borderRight string
}

func flipUpDown(t tile) tile {
	newTile := tile{
		ID:          t.ID,
		borderTop:   t.borderBot,
		borderBot:   t.borderTop,
		borderLeft:  reverseString(t.borderLeft),
		borderRight: reverseString(t.borderRight)}
	return newTile
}

func rotate(t tile) tile {
	newTile := tile{
		ID:          t.ID,
		borderTop:   t.borderLeft,
		borderBot:   t.borderRight,
		borderLeft:  t.borderBot,
		borderRight: t.borderTop}
	return newTile
}

func reverseString(s string) string {
	out := ""
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}

func contains(s string, list []string) bool {
	for _, el := range list {
		if s == el {
			return true
		}
	}
	return false
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isMatch(tile1 tile, tile2 tile) bool {
	t1 := []string{
		tile1.borderTop, tile1.borderBot, tile1.borderLeft, tile1.borderRight,
		reverseString(tile1.borderTop), reverseString(tile1.borderBot), reverseString(tile1.borderLeft), reverseString(tile1.borderRight),
	}
	t2 := []string{
		tile2.borderTop, tile2.borderBot, tile2.borderLeft, tile2.borderRight,
		reverseString(tile2.borderTop), reverseString(tile2.borderBot), reverseString(tile2.borderLeft), reverseString(tile2.borderRight),
	}
	for _, border := range t1 {
		if contains(border, t2) {
			return true
		}
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

	tiles := []tile{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Tile") {
			tileID, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[1], ":")[0])
			t := tile{
				ID:          tileID,
				borderLeft:  "",
				borderRight: "",
			}

			i := 0
			for scanner.Scan() {
				imgPart := scanner.Text()

				if i == 0 {
					t.borderTop = imgPart
				}
				if i == 9 {
					t.borderBot = imgPart
					t.borderLeft += string(imgPart[0])
					t.borderRight += string(imgPart[len(imgPart)-1])

					break
				}
				t.borderLeft += string(imgPart[0])
				t.borderRight += string(imgPart[len(imgPart)-1])

				i++
			}
			tiles = append(tiles, t)
		}
	}

	for _, t := range tiles {
		fmt.Println(t.ID, t.borderTop, t.borderBot, t.borderLeft, t.borderRight)
	}

	// part 1
	match := map[int][]int{}
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}
			if isMatch(tiles[i], tiles[j]) {
				match[tiles[i].ID] = append(match[tiles[i].ID], tiles[j].ID)
			}
		}
	}

	corners := 1
	for k, v := range match {
		matchingTilesCount := len(unique(v))
		if matchingTilesCount == 2 {
			corners *= k
		}
	}
	fmt.Println(corners)

	// TODO: part 2
	file.Close()
}
