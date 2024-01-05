package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re       *regexp.Regexp = regexp.MustCompile(`\d`)
	valueMap map[string]int = map[string]int{
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
)

func getNumbers(line string) ([][]int, []int) {
	var indexes [][]int
	var numbersInLine []int

	for key := range valueMap {
		if strings.Contains(line, key) {
			regex := regexp.MustCompile(key)
			allCoordinates := regex.FindAllStringIndex(line, -1)
			for _, v := range allCoordinates {
				indexes = append(indexes, v)
				numbersInLine = append(numbersInLine, valueMap[line[v[0]:v[1]]])

			}
		}
	}
	return indexes, numbersInLine
}

func getFirstAndLast(nums []int) [2]int {
	if len(nums) == 0 {
		log.Fatal("No Input Array")
	}
	if len(nums) == 1 {
		return [2]int{nums[0], nums[0]}
	}
	return [2]int{nums[0], nums[len(nums)-1]}
}

func main() {
	// Import files
	input, err := os.ReadFile("files/small-input-part2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	for _, line := range bytes.Split(input, []byte("\n")) {
		lines = append(lines, string(line[:]))
	}

	var numbers [][][]int
	// insert all digits that are already there
	for _, line := range lines {
		newNumber := re.FindAllStringIndex(line, -1)
		numbers = append(numbers, newNumber)
	}

	// Part 1:
	// sum1 := 0
	for i, line := range lines {
		var numbersInLine []int
		for _, v := range numbers[i] {
			n, err := strconv.Atoi(line[v[0]:v[1]])
			if err != nil {
				log.Fatal(err)
			}
			numbersInLine = append(numbersInLine, n)
		}

	}

	for i, line := range lines {
		fmt.Println(line)
		a, _ := getNumbers(line)
		fmt.Println(a)
		for _, value := range a {
			numbers[i] = append(numbers[i], value)
		}
		fmt.Println(numbers[i])
		// Add numbers to array
		// for i, stringIndex := range numbers[index] {
		// 	intsPerLine[i] = append(intsPerLine[i],
		// line[stringIndex[0]:stringIndex[1]])
		// }
		// fmt.Printf("%q\n", intsPerLine)
	}
	for i := range numbers {
		for _, w := range numbers[i] {
			fmt.Print(w, ": ")
			fmt.Print(lines[i][w[0]:w[1]], ";   ")
		}
		fmt.Println("")
	}
	// fmt.Println(sum)
}
