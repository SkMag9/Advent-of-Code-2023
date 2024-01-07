package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Read file at the provided input path
func readInput(path string) []byte {
	// Import files
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

// Parse File input and return slice of strings with each value being a line
// (separated by \n).
func getLines(input []byte) []string {
	var lines []string
	for _, line := range bytes.Split(input, []byte("\n")) {
		lines = append(lines, string(line[:]))
	}

	return lines
}

func getNumericDigits(line string) [][2]int {
	// Collection of [2]int values with:
	// numbers[i][0] = value
	// numbers[i][1] = starting index in line (used to determine position in line)
	var numbers [][2]int
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		number, err := strconv.Atoi(line[match[0]:match[1]])
		if err != nil {
			log.Fatal(err)
		}
		numberAndIndex := [2]int{number, match[0]}
		numbers = append(numbers, numberAndIndex)
	}

	return numbers
}

// Same as getNumericDigits() but for getting the digits from the words
func getWordDigits(line string) [][2]int {
	var numbers [][2]int
	valueMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for key, value := range valueMap {
		re := regexp.MustCompile(key)
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			numberAndIndex := [2]int{value, match[0]}
			numbers = append(numbers, numberAndIndex)
		}
	}

	return numbers
}

func getAllDigits(line string) [][2]int {
	numericDigits := getNumericDigits(line)
	wordDigits := getWordDigits(line)
	var allDigits [][2]int

	allDigits = append(allDigits, numericDigits...)
	allDigits = append(allDigits, wordDigits...)

	return allDigits
}

// Input: Slice of [2]int where numbers[i][0] is a digit and numbers[i][1] is
// the index in the lines.
// Retrurn: two digit int with first digit being the first digit and second
// being the last digit in a line.
func getInstruction(numbers [][2]int) (int, error) {
	if len(numbers) == 0 {
		err := errors.New("Input slice is empty!")
		return 0, err
	}
	var (
		highestIndexElement int = 0                // default: first index
		lowestIndexElement  int = len(numbers) - 1 // default: last index
	)
	for positionInSlice, number := range numbers {
		if number[1] > numbers[highestIndexElement][1] {
			highestIndexElement = positionInSlice
		}
		if number[1] < numbers[lowestIndexElement][1] {
			lowestIndexElement = positionInSlice
		}
	}

	firstDigit := numbers[lowestIndexElement][0]
	secondDigit := numbers[highestIndexElement][0]

	instructionString := fmt.Sprint(firstDigit) + fmt.Sprint(secondDigit)

	instruction, err := strconv.Atoi(instructionString)
	if err != nil {
		return 0, err
	}
	return instruction, nil
}

func part1(path string, fullFile bool) {
	var label string

	if fullFile {
		label = "Full File"
	} else {
		label = "Test File"
	}

	fmt.Printf("Part 1: %v\n", label)

	input := readInput(path)
	lines := getLines(input)

	var sum int = 0

	for _, line := range lines {
		digitsInLine := getNumericDigits(line)
		instruction, err := getInstruction(digitsInLine)
		if err == nil {
			sum += instruction
		}
		// Else: nothing since it throws errors if line has no digits
	}

	fmt.Printf("Result: %v\n", sum)
}

func part2(path string, fullFile bool) {
	var label string

	if fullFile {
		label = "Full File"
	} else {
		label = "Test File"
	}

	fmt.Printf("Part 2: %v\n", label)

	input := readInput(path)
	lines := getLines(input)

	var sum int = 0

	for _, line := range lines {
		digitsInLine := getAllDigits(line)
		instruction, err := getInstruction(digitsInLine)
		if err == nil {
			sum += instruction
		}
		// Else: nothing since it throws errors if line has no digits
	}

	fmt.Printf("Result: %v\n", sum)
}

func main() {
	part1("files/small-input-part1.txt", false)
	part2("files/small-input-part2.txt", false)

	part1("files/input.txt", true)
	part2("files/input.txt", true)
}
