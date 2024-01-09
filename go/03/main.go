package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readFile(path string) []byte {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func getLines(input []byte) []string {
	var lines []string
	for _, line := range bytes.Split(input, []byte("\n")) {
		lines = append(lines, string(line[:]))
	}
	return lines
}

func getAllNumbers(lines []string) [][][2]int {
	re := regexp.MustCompile(`\d+`)

	var schema [][][2]int = make([][][2]int, len(lines))

	for lineIndex, line := range lines {
		if line == "" {
			continue
		}

		var lineNumbers [][2]int
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			lineNumbers = append(lineNumbers, [2]int{match[0], match[1]})
		}

		schema[lineIndex] = lineNumbers
	}

	return schema
}

func filterPartNumbers(lines []string, numbers [][][2]int) []int {
	var partNumbers []int

	for lineIndex, numbersInLine := range numbers {
		for _, numberIndex := range numbersInLine {
			if isPartNumber(lines, lineIndex, numberIndex) {
				number, err := strconv.Atoi(
					lines[lineIndex][numberIndex[0]:numberIndex[1]],
				)
				if err != nil {
					log.Fatal(err)
				}

				partNumbers = append(partNumbers, number)
			}
		}
	}

	return partNumbers
}

func isSymbol(s string) bool {
	re := regexp.MustCompile(`\pS`)
	return re.MatchString(s)
}

func isPartNumber(lines []string, lineIndex int, numberIndex [2]int) bool {
	topLine := false
	bottomLine := false
	leftCol := false
	rightCol := false

	if lineIndex == 0 {
		topLine = true
	}
	if lineIndex == len(lines)-2 {
		bottomLine = true
	}
	if numberIndex[0] == 0 {
		leftCol = true
	}
	if numberIndex[1] == len(lines[0]) {
		rightCol = true
	}

	if isSymbol(lines[lineIndex][numberIndex[0]-1:numberIndex[0]]) || // left
		isSymbol(lines[lineIndex][numberIndex[1]:numberIndex[1]+1]) { // right
		return true
	}

	return false
}

func getPartNumberSum(lines []string) int {
	partNumberSum := 0

	allNumbers := getAllNumbers(lines)
	allNumbers = nil
	fmt.Println(allNumbers)

	return partNumberSum
}

func part1(path string, fullFile bool) {
	label := "Part 1:"

	if !fullFile {
		label = "Part 1: Test File"
	}

	fmt.Println(label)

	input := readFile(path)
	lines := getLines(input)

	fmt.Printf("Result: %v\n", getPartNumberSum(lines))
}

func part2(path string, fullFile bool) {
	label := "Part 2:"

	if !fullFile {
		label = "Part 2: Test File"
	}

	fmt.Println(label)

	// input := readFile(path)
	// lines := getLines(input)

	// fmt.Printf("Result: %v\n", getPower(lines))
}

func main() {
	part1("files/small-input.txt", false)
}
