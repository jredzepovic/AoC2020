package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	count int
	color string
}

func parseRule(line string) []rule {
	var rules []rule

	pattern := "((\\d+) (.*?)) bag"
	rgx := regexp.MustCompile(pattern)
	parsed := rgx.FindAllStringSubmatch(line, -1)

	for i := 0; i < len(parsed); i++ {
		cnt, _ := strconv.Atoi(parsed[i][2])
		clr := parsed[i][3]

		rules = append(rules, rule{count: cnt, color: clr})
	}

	return rules
}

func isValid(color string, rules []rule) bool {
	for _, r := range rules {
		if color == r.color {
			return true
		}
	}
	return false
}

func countBags(color string, luggageRules map[string][]rule) int {
	sum := 1

	for _, r := range luggageRules[color] {
		sum += r.count * countBags(r.color, luggageRules)
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	luggageRules := map[string][]rule{}
	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), " bags contain ")

		luggageRules[lineSplitted[0]] = append(luggageRules[lineSplitted[0]], parseRule(lineSplitted[1])...)
	}

	isFound := false
	next := []string{}
	current := []string{"shiny gold"}
	shinyGoldContainers := map[string]int{}
	for {
		for _, c := range current {
			for k, v := range luggageRules {
				if isValid(c, v) {
					isFound = true
					shinyGoldContainers[k]++
					next = append(next, k)
				}
			}
		}
		current = next
		next = nil
		if !isFound {
			break
		}
		isFound = false
	}

	// part 1
	fmt.Println(len(shinyGoldContainers))

	// part 2
	fmt.Println(countBags("shiny gold", luggageRules) - 1)
	file.Close()
}
