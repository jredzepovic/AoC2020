package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func getLocation(locations map[uint64]int) uint64 {
	max := 0
	maxVals := map[int]int{}
	id := uint64(0)

	for k, v := range locations {
		maxVals[v]++
		if v > max {
			max = v
			id = k
		}
	}
	if maxVals[max] == 1 {
		return id
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

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ranges := parseRanges(lines[0:20])

	// part 1
	intervalTree := augmentedtree.New(1)
	for _, r := range ranges {
		intervalTree.Add(augmentedtree.Interval(r))
	}

	notValid := int64(0)
	validTickets := [][]int64{}
	for i := 25; i < len(lines); i++ {
		t := parseTicket(lines[i])
		for _, n := range t {
			var k augmentedtree.Interval = interval{id: 0, low: n, high: n}
			if len(intervalTree.Query(k)) == 0 {
				notValid += n
			} else {
				validTickets = append(validTickets, t)
			}
		}
	}
	fmt.Println(notValid)

	// part 2
	myTicket := parseTicket(lines[22])
	locs := map[uint64]int{}

	i := 0
	found := 0
	for found < len(myTicket) {
		key := map[uint64]int{}

		for j := 0; j < len(validTickets); j++ {
			var k augmentedtree.Interval = interval{id: 0, low: validTickets[j][i], high: validTickets[j][i]}

			intervals := intervalTree.Query(k)
			for _, intv := range intervals {
				if _, exists := locs[uint64(math.Ceil(float64(intv.ID())/float64(2)))]; !exists {
					key[uint64(math.Ceil(float64(intv.ID())/float64(2)))]++
				}
			}
		}

		loc := getLocation(key)
		if loc != 0 {
			locs[loc] = i
			found++
		}

		i = (i + 1) % len(myTicket)
	}
	final := int64(1)
	for k, v := range locs {
		if k < 7 {
			final *= myTicket[v]
		}
	}
	fmt.Println(final)
}
