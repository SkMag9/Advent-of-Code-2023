package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`\d{1}`)

func goThorughArray(fn func([]string) int64, s []string) int64 {
	var sum int64 = 0
	for _, v := range s {
		numbers := re.FindAllString(v, -1)
		if len(numbers) != 0 {
			sum += fn(numbers)
		}
	}
	return sum

	// TODO:
	// [ ] addition to separate function - keep only iteration in this func
	// [ ] replacement operation
	// [ ] main() comments for part 1 and 2 goThorughArray(), then replace and
	// goThorughArray() again
}

//	func replaceNumbers(s []string) {
//		valueMap := map[string]string{
//			"one":   "1",
//			"two":   "2",
//			"three": "3",
//			"four":  "4",
//			"five":  "5",
//			"six":   "6",
//			"seven": "7",
//			"eight": "8",
//			"nine":  "9",
//		}
//
//		// Implement replacements
//	}
//

func stringToInt(s []string) int64 {
	calibrationValue, err := strconv.ParseInt(
		(s[0] + s[len(s)-1]),
		10,
		64,
	)
	if err != nil {
		log.Fatal(err)
	}
	return calibrationValue
}

func main() {
	inputFile, err := os.ReadFile("./files/small-input.txt")
	// inputFile, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions := strings.Fields(string(inputFile))

	println(goThorughArray(stringToInt, instructions))
}
