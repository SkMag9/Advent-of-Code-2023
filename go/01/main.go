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

func readInput(path string) []byte {
	// Import files
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

func getDigits(line string) [][2]int {
	// Collection of [2]int values with:
	// numbers[i][0] = value
	// numbers[i][1] = starting index in line (used to determine position in line)
	var numbers [][2]int
	re := regexp.MustCompile(`\d`)

	matches := re.FindAllStringIndex(line, -1)

	for _, v := range matches {
		number, err := strconv.Atoi(line[v[0]:v[1]])
		if err != nil {
			log.Fatal(err)
		}
		numberAndIndex := []int{number, v[0]}
		numbers = append(numbers, [2]int(numberAndIndex))
	}

	return numbers
}

// func getWordDigits(line string) [][2]int {
// 	// Same as getDigits() but for getting the digits from the words
// 	var numbers [][2]int
// 	valueMap := map[string]int{
// 		"zero":  0,
// 		"one":   1,
// 		"two":   2,
// 		"three": 3,
// 		"four":  4,
// 		"five":  5,
// 		"six":   6,
// 		"seven": 7,
// 		"eight": 8,
// 		"nine":  9,
// 	}
//
// 	// magic
//
// 	return numbers
// }

func getAllNumbers(line string) {
	getDigits(line)
	// getWordDigits(line)
}

func getFirstAndLast(nums [][2]int) (int, error) {
	if len(nums) == 0 {
		err := errors.New("Input slice is empty!")
		return 0, err
	}

	var (
		lowestIndexElement  int = 0             // default: first index
		highestIndexElement int = len(nums) - 1 // default: last index
	)
	for i, v := range nums {
		if v[1] == lowestIndexElement || v[1] == highestIndexElement {
			log.Fatal("Error in data processing: Two values with the same index!")
		}
		if v[1] > highestIndexElement {
			highestIndexElement = i
		}
		if v[1] < lowestIndexElement {
			lowestIndexElement = i
		}
	}

	firstDigit := nums[lowestIndexElement][0]
	secondDigit := nums[highestIndexElement][0]

	var instructionString string = string(firstDigit) + string(secondDigit)

	instruction, err := strconv.Atoi(instructionString)
	if err != nil {
		return 0, err
	}

	return instruction, nil
}

func main() {
	// Setup
	input := readInput("files/small-input-part2.txt")
	lines := getLines(input)
	fmt.Println(lines)

	// Part1
	var sumPart1 int = 0
	for _, line := range lines {
		digitsInLine := getDigits(line)
		instruction, err := getFirstAndLast(digitsInLine)
		if err == nil {
		}
		// Else: nothing -> Errors if line has no digits
	}
}

// var numbers [][][]int
// // insert all digits that are already there
// for _, line := range lines {
// 	newNumber := re.FindAllStringIndex(line, -1)
// 	numbers = append(numbers, newNumber)
// }
//
// // Part 1:
// // sum1 := 0
// for i, line := range lines {
// 	var numbersInLine []int
// 	for _, v := range numbers[i] {
// 		n, err := strconv.Atoi(line[v[0]:v[1]])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		numbersInLine = append(numbersInLine, n)
// 	}
// 	fmt.Println("nil:", numbersInLine)
// }
//
// for i, line := range lines {
// 	fmt.Println(line)
// 	a, _ := getNumbers2(line)
// 	fmt.Println(a)
// 	for _, value := range a {
// 		numbers[i] = append(numbers[i], value)
// 	}
// 	fmt.Println(numbers[i])
// 	// Add numbers to array
// 	// for i, stringIndex := range numbers[index] {
// 	// 	intsPerLine[i] = append(intsPerLine[i],
// 	// line[stringIndex[0]:stringIndex[1]])
// 	// }
// 	// fmt.Printf("%q\n", intsPerLine)
// }
// for i := range numbers {
// 	for _, w := range numbers[i] {
// 		fmt.Print(w, ": ")
// 		fmt.Print(lines[i][w[0]:w[1]], ";   ")
// 	}
// 	fmt.Println("")
// }
// // fmt.Println(sum)

// func getNumbers2(line string) ([][]int, []int) {
// 	var indexes [][]int
// 	var numbersInLine []int
//
// 	for key := range valueMap {
// 		if strings.Contains(line, key) {
// 			regex := regexp.MustCompile(key)
// 			allCoordinates := regex.FindAllStringIndex(line, -1)
// 			for _, v := range allCoordinates {
// 				indexes = append(indexes, v)
// 				numbersInLine = append(numbersInLine, valueMap[line[v[0]:v[1]]])
//
// 			}
// 		}
// 	}
// 	return indexes, numbersInLine
// }
