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
	re := regexp.MustCompile(`[^\.]`)
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
	skipBottom := bool(lineIndex >= len(lines)-2) // last value always empty
	skipLeft := bool(numberIndex[0] == 0)
	skipRight := bool(numberIndex[1] == len(lines[0]))

	if !(skipTop || skipLeft) {
		if isSymbol(lines[topLeft[0]][topLeft[1]:topLeft[2]]) {
			return true
		}
	}
	if !(skipTop || skipRight) {
		if isSymbol(lines[topRight[0]][topRight[1]:topRight[2]]) {
			return true
		}
	}
	if !(skipBottom || skipLeft) {
		if isSymbol(lines[downLeft[0]][downLeft[1]:downLeft[2]]) {
			return true
		}
	}
	if !(skipBottom || skipRight) {
		if isSymbol(lines[downRight[0]][downRight[1]:downRight[2]]) {
			return true
		}
	}
	if !skipTop {
		if isSymbol(lines[up[0]][up[1]:up[2]]) {
			return true
		}
	}
	if !skipBottom {
		if isSymbol(lines[down[0]][down[1]:down[2]]) {
			return true
		}
	}
	if !skipLeft {
		if isSymbol(lines[left[0]][left[1]:left[2]]) {
			return true
		}
	}
	if !skipRight {
		if isSymbol(lines[right[0]][right[1]:right[2]]) {
			return true
		}
	}
	return false
}

func getPartNumberSum(lines []string) int {
	partNumberSum := 0

	allNumbers := getAllNumbers(lines)

	for lineIndex, line := range allNumbers {
		for _, number := range line {
			if isPartNumber(lines, lineIndex, number) {
				partNumber, err := strconv.Atoi(lines[lineIndex][number[0]:number[1]])
				if err != nil {
					log.Fatal(err)
				}

				partNumberSum += partNumber
			}
		}
	}

	return partNumberSum
}

type neighbours struct {
	topLeft     bool
	topRight    bool
	bottomLeft  bool
	bottomRight bool
	top         bool
	bottom      bool
	left        bool
	right       bool
}

func getGears(line string) [][2]int {
	var gears [][2]int
	re := regexp.MustCompile(`\*`)
	matches := re.FindAllStringIndex(line, -1)
	for _, match := range matches {
		gears = append(gears, [2]int{match[0], match[1]})
	}
	return gears
}

func getGearNeighbours(
	lines []string,
	lineIndex int,
	gearPositions [][2]int,
) neighbours {
	var gearNeighbours neighbours

	for _, gear := range gearPositions {
		if lineIndex != 0 {
			gearNeighbours.topLeft = true
			gearNeighbours.topRight = true
			gearNeighbours.top = true
		}

		if lineIndex != len(lines)-2 {
			gearNeighbours.bottomLeft = true
			gearNeighbours.bottomRight = true
			gearNeighbours.bottom = true
		}

		if gear[0] == 0 {
			gearNeighbours.topLeft = false
			gearNeighbours.bottomLeft = false
			gearNeighbours.left = false
		}

		if gear[1] == len(lines[0]) {
			gearNeighbours.topRight = false
			gearNeighbours.bottomRight = false
			gearNeighbours.right = false

		} else {
			gearNeighbours.topRight = false
			gearNeighbours.bottomRight = false
			gearNeighbours.right = false
		}

		fmt.Println(gearNeighbours)
	}

	return gearNeighbours
}

func getGearNumbers(lines []string, lineIndex int, gearLocation [2]int) []int {
	return []int{0}
}

func getGearRatioSum(lines []string) int {
	gearRatioSum := 0

	for lineIndex, line := range lines {
		gearsInLine := getGears(line)
		for _, gear := range gearsInLine {
			fmt.Println(lineIndex, gear)
		}
	}

	return gearRatioSum
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

	input := readFile(path)
	lines := getLines(input)

	fmt.Printf("Result: %v\n", getGearRatioSum(lines))
}

func main() {
	part1("files/small-input.txt", false)
	part1("files/input.txt", true)
	part2("files/small-input.txt", false)
}
