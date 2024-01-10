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

// 0: lineIndex
// 1: start index
// 2: end index
func getNeighbours(
	lineIndex int, numberIndex [2]int,
) (up, down, left, right, topLeft, topRight, downLeft, downRight [3]int) {
	up = [3]int{lineIndex - 1, numberIndex[0], numberIndex[1]}
	down = [3]int{lineIndex + 1, numberIndex[0], numberIndex[1]}
	left = [3]int{lineIndex, numberIndex[0] - 1, numberIndex[0]}
	right = [3]int{lineIndex, numberIndex[1], numberIndex[1] + 1}
	topLeft = [3]int{lineIndex - 1, numberIndex[0] - 1, numberIndex[0]}
	topRight = [3]int{lineIndex - 1, numberIndex[1], numberIndex[1] + 1}
	downLeft = [3]int{lineIndex + 1, numberIndex[0] - 1, numberIndex[0]}
	downRight = [3]int{lineIndex + 1, numberIndex[1], numberIndex[1] + 1}

	return
}

func isPartNumber(lines []string, lineIndex int, numberIndex [2]int) bool {
	up, down, left, right, topLeft, topRight, downLeft, downRight := getNeighbours(
		lineIndex,
		numberIndex,
	)
	skipTop := bool(lineIndex == 0)
	skipBottom := bool(lineIndex == len(lines)-2) // last value always empty
	skipLeft := bool(numberIndex[0] == 0)
	skipRight := bool(numberIndex[1] == len(lines[0]))

	fmt.Println(
		"Line: ",
		lines[lineIndex],
		"Bools: ",
		skipTop,
		skipBottom,
		skipLeft,
		skipRight,
	)

	fmt.Println(
		"Checks:",
		!skipTop && !skipLeft,
		!skipTop && !skipRight,
		!skipBottom && !skipLeft,
		!skipBottom && !skipRight,
		!skipTop,
		!skipBottom,
		!skipLeft,
		!skipRight,
	)

	if !skipTop && !skipLeft {
		fmt.Println(lines[topLeft[0]][topLeft[1]:topLeft[2]])
		if isSymbol(lines[topLeft[0]][topLeft[1]:topLeft[2]]) {
			return true
		}
	}
	if !skipTop && !skipRight {
		fmt.Println(lines[topRight[0]][topRight[1]:topRight[2]])
		if isSymbol(lines[topRight[0]][topRight[1]:topRight[2]]) {
			return true
		}
	}
	if !skipBottom && !skipLeft {
		fmt.Println(lines[downLeft[0]][downLeft[1]:downLeft[2]])
		if isSymbol(lines[downLeft[0]][downLeft[1]:downLeft[2]]) {
			return true
		}
	}
	if !skipBottom && !skipRight {
		fmt.Println(lines[downRight[0]][downRight[1]:downRight[2]])
		if isSymbol(lines[downRight[0]][downRight[1]:downRight[2]]) {
			return true
		}
	}
	if !skipTop {
		fmt.Println(lines[up[0]][up[1]:up[2]])
		if isSymbol(lines[up[0]][up[1]:up[2]]) {
			return true
		}
	}
	if !skipBottom {
		fmt.Println(lines[down[0]][down[1]:down[2]])
		if isSymbol(lines[down[0]][down[1]:down[2]]) {
			return true
		}
	}
	if !skipLeft {
		fmt.Println(lines[left[0]][left[1]:left[2]])
		if isSymbol(lines[left[0]][left[1]:left[2]]) {
			return true
		}
	}
	if !skipRight {
		fmt.Println(lines[right[0]][right[1]:right[2]])
		if isSymbol(lines[right[0]][right[1]:right[2]]) {
			return true
		}
	}
	return false
}

func getPartNumberSum(lines []string) int {
	partNumberSum := 0

	allNumbers := getAllNumbers(lines)
	fmt.Println(allNumbers)

	for lineIndex, line := range allNumbers {
		for _, number := range line {
			fmt.Println(lines[lineIndex][number[0]:number[1]])
			if isPartNumber(lines, lineIndex, number) {
				partNumber, err := strconv.Atoi(lines[lineIndex][number[0]:number[1]])
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(partNumber)
				partNumberSum += partNumber
			} else {
				fmt.Println(line, lineIndex, number)
				fmt.Println("Not a Part: ", isPartNumber(lines, lineIndex, number))
			}
		}
	}

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
