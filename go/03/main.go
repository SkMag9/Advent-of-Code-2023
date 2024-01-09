package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
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

func getAllNumbers(lines []string) [][][2]string {
	re := regexp.MustCompile(`\d+`)

	var schema [][][2]string

	for lineIndex, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindAllStringIndex(line, -1)

	}

	return schema
}

func isAdjacent() bool {
	return true
}

func getPartNumberSum(lines []string) int {
	partNumberSum := 0

	fmt.Printf("%q\n", lines)
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
