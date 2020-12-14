package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseWriteInstruction(line string) (int64, int64) {
	var address, value int64

	pattern := "(\\d+)"
	rgx := regexp.MustCompile(pattern)

	splitted := strings.Split(line, " ")
	value, _ = strconv.ParseInt(splitted[2], 10, 64)

	parsed := rgx.FindStringSubmatch(splitted[0])
	address, _ = strconv.ParseInt(parsed[1], 10, 64)

	return address, value
}

func applyMaskV1(value int64, mask string) int64 {
	updated := ""
	valueBinary := strconv.FormatInt(value, 2)
	valueBinary = strings.Repeat("0", len(mask)-len(valueBinary)) + valueBinary

	for i := 0; i < len(mask); i++ {
		if mask[i] != 'X' {
			updated += string(mask[i])
		} else {
			updated += string(valueBinary[i])
		}
	}

	newVal, _ := strconv.ParseInt(updated, 2, 64)
	return newVal
}

func applyMaskV2(value int64, mask string) string {
	updated := ""
	valueBinary := strconv.FormatInt(value, 2)
	valueBinary = strings.Repeat("0", len(mask)-len(valueBinary)) + valueBinary

	for i := 0; i < len(mask); i++ {
		if mask[i] == '0' {
			updated += string(valueBinary[i])
		} else if mask[i] == '1' {
			updated += string('1')
		} else {
			updated += string('X')
		}
	}

	return updated
}

func expandAddress(address string) []int64 {
	addresses := []int64{}

	countX := strings.Count(address, "X")

	if countX == 0 {
		newVal, _ := strconv.ParseInt(address, 2, 64)
		return []int64{newVal}
	}

	for i := int64(0); i < int64(math.Pow(float64(2), float64(countX))); i++ {
		valueBinary := strconv.FormatInt(i, 2)
		valueBinary = strings.Repeat("0", countX-len(valueBinary)) + valueBinary

		updated := strings.Split(address, "")
		k := 0
		for j := 0; j < len(updated); j++ {
			if k == countX {
				break
			}
			if address[j] == 'X' {
				updated[j] = string(valueBinary[k])
				k++
			}
		}

		newVal, _ := strconv.ParseInt(strings.Join(updated, ""), 2, 64)
		addresses = append(addresses, newVal)
	}

	return addresses
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

	// part 1
	var mask string
	memory := map[int64]int64{}
	for _, l := range lines {
		if strings.Contains(l, "mask") {
			mask = strings.Split(l, " ")[2]
		} else {
			address, value := parseWriteInstruction(l)
			value = applyMaskV1(value, mask)

			memory[address] = value
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}

	fmt.Println(sum)

	// part 2
	memory = map[int64]int64{}
	for _, l := range lines {
		if strings.Contains(l, "mask") {
			mask = strings.Split(l, " ")[2]
		} else {
			address, value := parseWriteInstruction(l)
			decodedAddress := applyMaskV2(address, mask)

			for _, a := range expandAddress(decodedAddress) {
				memory[a] = value
			}
		}
	}

	sum = int64(0)
	for _, v := range memory {
		sum += v
	}

	fmt.Println(sum)
	file.Close()
}
