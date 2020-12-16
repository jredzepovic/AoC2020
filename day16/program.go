package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Workiva/go-datastructures/augmentedtree"
)

type interval struct {
	id   uint64
	low  int64
	high int64
}

func (i interval) LowAtDimension(dim uint64) int64 {
	return i.low
}

func (i interval) HighAtDimension(dim uint64) int64 {
	return i.high
}

func (i interval) OverlapsAtDimension(other augmentedtree.Interval, dim uint64) bool {
	return i.LowAtDimension(dim) <= other.HighAtDimension(dim) && other.LowAtDimension(dim) <= i.HighAtDimension(dim)
}

func (i interval) ID() uint64 {
	return i.id
}

func parseRanges(lines []string) []interval {
	ints := []interval{}

	for i, l := range lines {
		splitted := strings.Split(l, " ")

		for j, r := range []string{splitted[len(splitted)-3], splitted[len(splitted)-1]} {
			low, _ := strconv.ParseInt(strings.Split(r, "-")[0], 10, 64)
			high, _ := strconv.ParseInt(strings.Split(r, "-")[1], 10, 64)

			ints = append(ints, interval{
				id:   uint64((i+1)*2 - j),
				low:  low,
				high: high,
			})
		}
	}

	return ints
}

func parseTicket(line string) []int64 {
	nums := []int64{}
	splitted := strings.Split(line, ",")

	for _, s := range splitted {
		n, _ := strconv.ParseInt(s, 10, 64)
		nums = append(nums, n)
	}

	return nums
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ranges := parseRanges(lines[0:20])

	//part 1
	intervalTree := augmentedtree.New(1)
	for _, r := range ranges {
		intervalTree.Add(augmentedtree.Interval(r))
	}

	notValid := int64(0)
	for i := 25; i < len(lines); i++ {
		for _, n := range parseTicket(lines[i]) {
			var k augmentedtree.Interval = interval{id: 0, low: n, high: n}
			if len(intervalTree.Query(k)) == 0 {
				notValid += n
			}
		}
	}
	fmt.Println(notValid)
}
