package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"regexp"
	"strconv"
)

func parseInstruction(instruction string) (string, complex128) {
	pattern := "(\\w+?)(\\d+)"

	rgx := regexp.MustCompile(pattern)
	parsed := rgx.FindStringSubmatch(instruction)
	val, _ := strconv.Atoi(parsed[2])

	return parsed[1], complex(float64(val), 0)
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

	directions := map[string]complex128{"N": 1i, "E": 1, "S": -1i, "W": -1}
	rotations := map[string]complex128{"L": 1i, "R": -1i}

	// part 1
	location := complex128(0)
	direction := directions["E"]
	for _, l := range lines {
		action, value := parseInstruction(l)

		if action == "F" {
			location += direction * value
		} else if action == "L" || action == "R" {
			direction *= cmplx.Pow(rotations[action], value/90)
		} else {
			location += directions[action] * value
		}
	}

	fmt.Println(math.Abs(real(location)) + math.Abs(imag(location)))

	// part 2
	location = complex128(0)
	waypoint := complex(10, 1)
	for _, l := range lines {
		action, value := parseInstruction(l)

		if action == "F" {
			location += waypoint * value
		} else if action == "L" || action == "R" {
			waypoint *= cmplx.Pow(rotations[action], value/90)
		} else {
			waypoint += directions[action] * value
		}
	}

	fmt.Println(math.Abs(real(location)) + math.Abs(imag(location)))
	file.Close()
}
