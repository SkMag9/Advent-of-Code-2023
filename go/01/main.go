package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions := strings.Fields(string(inputFile))
	re := regexp.MustCompile(`\d{1}`)

	var sum int64

	for _, v := range instructions {
		numbers := re.FindAllString(v, -1)
		if len(numbers) == 0 {
			log.Fatal("No Input")
		}
		calibrationValue, err := strconv.ParseInt(
			(numbers[0] + numbers[len(numbers)-1]),
			10,
			64,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			sum += calibrationValue
		}
	}

	fmt.Println(sum)
}
