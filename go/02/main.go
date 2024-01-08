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

func hasSingleColon(line string) bool {
	re := regexp.MustCompile(`:`)
	matches := re.FindAllString(line, -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func splitPrefixAndData(line string) (prefix, data string) {
	if !hasSingleColon(line) {
		log.Fatal("Number of colons isn't 1 for line:", line)
	}
	split := strings.Split(line, ":")

	prefix = split[0]
	data = split[1]

	return prefix, data
}

func getGameId(prefix string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(prefix)

	id, err := strconv.Atoi(match)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func splitDraws(data string) []string {
	return strings.Split(data, ";")
}

func splitColorsInDraw(draw string) [3][]int {
	drawValues := strings.Split(draw, ",")
	var colorValues [3][]int

	for _, drawValue := range drawValues {
		colorValue := strings.Split(drawValue, " ")

		amount, err := strconv.Atoi(colorValue[1])
		if err != nil {
			log.Fatal(err)
		}

		if colorValue[2] == "red" {
			colorValues[0] = append(colorValues[0], amount)
		}
		if colorValue[2] == "green" {
			colorValues[1] = append(colorValues[1], amount)
		}
		if colorValue[2] == "blue" {
			colorValues[2] = append(colorValues[2], amount)
		}
	}

	return colorValues
}

// Returns a slice of integers from the game for each color
// 0: red, 1: green, 2: blue
func getColorValues(data string) [3][]int {
	var colorValues [3][]int
	draws := splitDraws(data)

	for _, draw := range draws {
		for i := 0; i <= 2; i++ {
			colorValues[i] = append(colorValues[i], splitColorsInDraw(draw)[i]...)
		}
	}

	return colorValues
}

func isGamePossible(colorValues [3][]int, maxColorValues [3]int) bool {
	for colorIndex, color := range colorValues {
		for _, colorValue := range color {
			if colorValue > maxColorValues[colorIndex] {
				return false
			}
		}
	}
	return true
}

func getPossibleGameIdSum(lines []string, maxColorValues [3]int) int {
	var sum int = 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		prefix, data := splitPrefixAndData(line)

		colorValues := getColorValues(data)

		if isGamePossible(colorValues, maxColorValues) {
			gameId, err := getGameId(prefix)
			if err != nil {
				log.Fatal(err)
			}

			sum += gameId
		}
	}

	return sum
}

func part1(path string, fullFile bool, maxColorValues [3]int) {
	label := "Part 1:"

	if !fullFile {
		label = "Part 1: Test File"
	}

	fmt.Println(label)

	input := readFile(path)
	lines := getLines(input)

	fmt.Printf("Result: %v\n", getPossibleGameIdSum(lines, maxColorValues))
}

func main() {
	part1MaxColorValues := [3]int{
		12, // red
		13, // green
		14, // blue
	}
	part1("files/small-input-part1.txt", false, part1MaxColorValues)
	part1("files/input.txt", true, part1MaxColorValues)
}
