package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	position  int
	operation string
	argument  int
}

func parseInstruction(line string, position int) instruction {
	lineSplitted := strings.Split(line, " ")
	arg, _ := strconv.Atoi(lineSplitted[1])

	return instruction{position: position, operation: lineSplitted[0], argument: arg}
}

func executeInstruction(instruction instruction, accumulator int) (int, int) {
	if instruction.operation == "acc" {
		return executeJump(instruction.position, 1), executeAritmetic(instruction.argument, accumulator)
	}
	if instruction.operation == "jmp" {
		return executeJump(instruction.position, instruction.argument), accumulator
	}
	return executeJump(instruction.position, 1), accumulator
}

func executeJump(position int, offset int) int {
	return position + offset
}

func executeAritmetic(argument int, accumulator int) int {
	return accumulator + argument
}

func changeNopJmpOperation(position int, instructions []instruction) []instruction {
	updatedInstructions := make([]instruction, len(instructions))
	copy(updatedInstructions, instructions)

	if instructions[position].operation == "nop" {
		updatedInstructions[position].operation = "jmp"
	}
	if instructions[position].operation == "jmp" {
		updatedInstructions[position].operation = "nop"
	}
	return updatedInstructions
}

func runCode(instructions []instruction) (bool, int) {
	i := 0
	accumulator := 0
	executed := map[int]int{}
	for {
		if _, exists := executed[i]; exists {
			return false, accumulator
		}
		if i == len(instructions)-1 {
			return true, accumulator
		}
		executed[i]++

		i, accumulator = executeInstruction(instructions[i], accumulator)
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	position := 0
	instructions := []instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text(), position))
		position++
	}

	// part 1
	fmt.Println(runCode(instructions))

	// part 2
	nopJmpIndex := []int{}
	for i, ins := range instructions {
		if ins.operation == "nop" || ins.operation == "jmp" {
			nopJmpIndex = append(nopJmpIndex, i)
		}
	}

	for _, ind := range nopJmpIndex {
		updatedInstructions := changeNopJmpOperation(ind, instructions)

		success, acc := runCode(updatedInstructions)
		if success {
			fmt.Println(acc)
			break
		}
	}
	file.Close()
}
