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

type validate func(string) bool

func byrRule(value string) bool {
	val, err := strconv.Atoi(value)

	if err != nil || val < 1920 || val > 2002 {
		return false
	}
	return true
}

func iyrRule(value string) bool {
	val, err := strconv.Atoi(value)

	if err != nil || val < 2010 || val > 2020 {
		return false
	}
	return true
}

func eyrRule(value string) bool {
	val, err := strconv.Atoi(value)

	if err != nil || val < 2020 || val > 2030 {
		return false
	}
	return true
}

func hgtRule(value string) bool {
	pattern := "(\\d+)(cm|in)"

	result, err := regexp.MatchString(pattern, value)
	if !result || err != nil {
		return false
	}

	rgx := regexp.MustCompile(pattern)
	parsed := rgx.FindStringSubmatch(value)

	unit := parsed[2]
	hgt, err := strconv.Atoi(parsed[1])
	if err != nil {
		return false
	}

	if unit == "in" {
		if hgt >= 59 && hgt <= 76 {
			return true
		}
	} else if unit == "cm" {
		if hgt >= 150 && hgt <= 193 {
			return true
		}
	} else {
		return false
	}
	return false
}

func hclRule(value string) bool {
	pattern := "^#[0-9a-f]{6}$"
	result, err := regexp.MatchString(pattern, value)

	if !result || err != nil {
		return false
	}
	return true
}

func eclRule(value string) bool {
	switch value {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		return true
	}
	return false
}

func pidRule(value string) bool {
	if len(value) == 9 {
		if _, err := strconv.Atoi(value); err == nil {
			return true
		}
	}
	return false
}

func validPassport(passport string) int {
	var requiredKeys = map[string]validate{
		"byr": byrRule,
		"iyr": iyrRule,
		"eyr": eyrRule,
		"hgt": hgtRule,
		"hcl": hclRule,
		"ecl": eclRule,
		"pid": pidRule}

	for k, v := range requiredKeys {
		if strings.Contains(passport, k+":") {
			rgx := regexp.MustCompile(k + ":(\\S*)")
			passportValue := rgx.FindStringSubmatch(passport)[1]
			if !v(passportValue) {
				return 0
			}
		} else {
			return 0
		}
	}

	return 1
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passport []string
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			counter += validPassport(strings.Join(passport, " "))
			passport = nil
		} else {
			passport = append(passport, line)
		}
	}
	counter += validPassport(strings.Join(passport, " "))

	fmt.Println(counter)
	file.Close()
}
